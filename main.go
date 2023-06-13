package main

import (
	"context"
	"fmt"
	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func doMain(ctx context.Context) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := func() error {
			q := r.URL.Query()
			getQueryOrEnv := func(kQuery, kEnv string) string {
				if v := q.Get(kQuery); v != "" {
					return v
				}
				return os.Getenv(kEnv)
			}

			SDKSecretKey := getQueryOrEnv("secret", "TRTC_SECRETKEY")
			if SDKSecretKey == "" {
				return fmt.Errorf("env TRTC_SECRETKEY or query secret must be set")
			}

			sSDKAppID := getQueryOrEnv("appid", "TRTC_APPID")
			if sSDKAppID == "" {
				return fmt.Errorf("env TRTC_APPID or query appid must be set")
			}

			var SDKAppID int
			if v, err := strconv.ParseInt(sSDKAppID, 10, 64); err != nil || v <= 0 {
				return fmt.Errorf("TRTC_APPID=%v must be a positive integer", sSDKAppID)
			} else {
				SDKAppID = int(v)
			}

			tmpl, err := template.ParseFiles("tmpl/index.html")
			if err != nil {
				return err
			}

			streamID := q.Get("stream")
			if streamID == "" {
				streamID = fmt.Sprintf("%x", rand.Uint64())
				streamID = streamID[len(streamID)-7:]
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

			return tmpl.Execute(w, &struct {
				SDKAppID     int
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
			}{
				SDKAppID:        SDKAppID,
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
