package main

import (
	"context"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
)

type TmplResponse struct {
	Ready bool
	// It should be an integer, but we use string for HTML input.
	SDKAppID string
	// The secret key for TRTC, use mask for HTML input.
	SDKSecretKey string
	// The stream ID for WHIP, or room ID for TRTC.
	StreamID string
	// The WHIP endpoing.
	WHIPServer string
	// The session ID for WHIP, or user ID for TRTC.
	WHIPSessionID string
	// The Bearer Token for WHIP, or userSig for TRTC.
	WHIPBearerToken string
	// The WHEP endpoing.
	WHEPServer string
	// The session ID for WHEP, or user ID for TRTC.
	WHEPSessionID string
	// The Bearer Token for WHEP, or userSig for TRTC.
	WHEPBearerToken string
	// The TRTC user ID to pull stream.
	TRTCUserID string
	// The TRTC user signature to pull stream.
	TRTCUserSig string
}

func doMain(ctx context.Context) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := func() error {
			getQueryOrEnv := func(q url.Values, kQuery, kEnv string) string {
				if v := q.Get(kQuery); v != "" {
					return v
				}
				return os.Getenv(kEnv)
			}

			tmpl, err := template.ParseFiles("tmpl/index.html")
			if err != nil {
				return err
			}

			q := r.URL.Query()
			streamID := q.Get("stream")
			if streamID == "" {
				streamID = fmt.Sprintf("%x", rand.Uint64())
				streamID = streamID[len(streamID)-7:]
			}

			SDKSecretKey := getQueryOrEnv(q, "secret", "TRTC_SECRETKEY")
			sSDKAppID := getQueryOrEnv(q, "appid", "TRTC_APPID")
			ready := sSDKAppID != "" && SDKSecretKey != ""
			if !ready {
				return tmpl.Execute(w, &TmplResponse{})
			}

			var SDKAppID int
			if v, err := strconv.ParseInt(sSDKAppID, 10, 64); err != nil || v <= 0 {
				return fmt.Errorf("TRTC_APPID=%v must be a positive integer", sSDKAppID)
			} else {
				SDKAppID = int(v)
			}

			whipSessionID := fmt.Sprintf("%x", rand.Uint64())
			whipSessionID = whipSessionID[len(whipSessionID)-9:]
			whipUserSig, err := tencentyun.GenUserSig(SDKAppID, SDKSecretKey, whipSessionID, 3600)
			if err != nil {
				return err
			}
			whipServer := fmt.Sprintf("https://signaling.rtc.tencentcloud.com/v2/pub/%v/%v?SessionID=%v",
				SDKAppID, streamID, whipSessionID)

			whepSessionID := fmt.Sprintf("%x", rand.Uint64())
			whepSessionID = whepSessionID[len(whepSessionID)-9:]
			whepUserSig, err := tencentyun.GenUserSig(SDKAppID, SDKSecretKey, whepSessionID, 3600)
			if err != nil {
				return err
			}
			whepServer := fmt.Sprintf("https://signaling.rtc.tencentcloud.com/v2/sub/%v/%v?SessionID=%v",
				SDKAppID, streamID, whepSessionID)

			trtcUserID := fmt.Sprintf("%x", rand.Uint64())
			trtcUserID = trtcUserID[len(trtcUserID)-9:]
			trtcUserSig, err := tencentyun.GenUserSig(SDKAppID, SDKSecretKey, trtcUserID, 3600)
			if err != nil {
				return err
			}

			return tmpl.Execute(w, &TmplResponse{
				Ready:           true,
				SDKAppID:        sSDKAppID,
				SDKSecretKey:    strings.Repeat("*", len(SDKSecretKey)),
				StreamID:        streamID,
				WHIPServer:      whipServer,
				WHIPSessionID:   whipSessionID,
				WHIPBearerToken: whipUserSig,
				WHEPServer:      whepServer,
				WHEPSessionID:   whepSessionID,
				WHEPBearerToken: whepUserSig,
				TRTCUserID:      trtcUserID,
				TRTCUserSig:     trtcUserSig,
			})
		}(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	listen := os.Getenv("LISTEN")
	if listen == "" {
		listen = ":9000"
	}
	if !strings.Contains(listen, ":") {
		listen = ":" + listen
	}
	fmt.Println("listen on", listen, "please open", fmt.Sprintf("http://localhost%v", listen))

	return http.ListenAndServe(listen, nil)
}

func main() {
	if err := doMain(context.Background()); err != nil {
		fmt.Println(err.Error())
	}
}
