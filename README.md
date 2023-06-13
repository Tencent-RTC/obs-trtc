# OBS-TRTC

[![](https://img.shields.io/twitter/follow/TencentRTC?style=social)](https://twitter.com/TencentRTC)
[![](https://img.shields.io/badge/TencentRTC-YouTube-red)](https://www.youtube.com/@TencentRTC)
[![](https://badgen.net/discord/members/vDHty6ddrZ)](https://discord.gg/vDHty6ddrZ)

Demo for using OBS WHIP to publish a stream to the TRTC(Tencent Real-Time Communication) service.

## Overview

OBS features WHIP support, which opens up a wide range of fascinating use cases by 
harnessing the capabilities of both OBS and WHIP.

Since WHIP is a standardized protocol, it allows you to utilize HTML5 and various 
clients for publishing and playing real-time streams. Additionally, you can employ 
open-source tools to construct a real-time streaming platform.

TRTC(Tencent Real-Time Communication) cloud services are also available for utilizing 
a streaming platform with OBS WHIP support, particularly when you prefer not to create 
your own streaming platform or seek a more stable and dependable platform with dedicated 
support.

Additionally, TRTC(Tencent Real-Time Communication) offers a free trial with a 
certain number of streaming minutes, making it very easy to use.

If you require assistance or encounter any difficulties, please feel free to reach out 
to us via [Discord](https://discord.gg/vDHty6ddrZ).

## Prerequisites

Please make sure that you have the following prerequisites before proceeding:

- OBS with WHIP support, please downlaod from [OBS](https://obsproject.com/)
- TRTC(Tencent Real-Time Communication) account, please register at [here](https://trtc.tencentcloud.com/)

> Note: Currently, OBS WHIP has been merged into the master branch but has not been released yet. 
> You can download it from [here](https://github.com/obsproject/obs-studio/actions/runs/5227109208?pr=7926).

Next, you need to create a TRTC application and generate a bearer token for WHIP.

## Step 1: Create a TRTC application

Please follow the steps below to create a TRTC application:

1. Log in to the [TRTC console](https://console.cloud.tencent.com/trtc) and click **Application Management** on the left sidebar.
2. Click **Create Application** and enter the application name.
3. Click **Create** to create the application.

![001](https://github.com/Tencent-RTC/obs-trtc/assets/2777660/dce31494-ac4f-4844-b437-de3d244af678)

After the application is created, you can view the SDKAppID and SDKSecretKey on the application details page.

* SDKAppID: `2000xxxx`
* SDKSecretKey: `yyyyyy`

![003](https://github.com/Tencent-RTC/obs-trtc/assets/2777660/f54ad443-645a-4089-8dc7-5a34af48a335)

Following that, you must generate a bearer token for WHIP, which will be 
utilized in OBS.

## Step 2: Create a Bearer Token for WHIP

You can directly visit https://trtc.ossrs.io/obs-trtc/?appid=2000xxxx&secret=yyyyyy
to create a WHIP bearer token. Ensure that you replace the appid with your own `SDKAppID` and 
secret with your own `SDKSecretKey`.

In the `OBS WHIP` section, you will find the generated WHIP Bearer Token for configuring OBS.

![004](https://github.com/Tencent-RTC/obs-trtc/assets/2777660/29f1ac74-cb3a-4c6e-a3e4-e6c92ea2aab0)

The configuration settings for OBS WHIP are as follows:

* Service: `WHIP`
* Server: `https://signaling.rtc.tencentcloud.com/v2/pub/2000xxxx/xxxxxxx?SessionID=xxxxxxxxx`
* Bearer Token: `xxxxxx-yyyyyy-zzzzzz`

Alternatively, you can operate your own server to produce the token by utilizing 
the following command line:

```bash
git clone https://github.com/Tencent-RTC/obs-trtc.git
cd obs-trtc
TRTC_APPID=2000xxxx TRTC_SECRETKEY=yyyyyy go run .
```

Next, navigate to the webpage http://localhost:9000/ to generate a WHIP bearer token.

## Step 3: Configure OBS

Please follow the steps below to configure OBS:

1. Open OBS and click **Settings**.
2. Click **Stream** on the left sidebar.
3. Select **WHIP** for **Service**.
4. Enter the **Server** and **Bearer Token** generated in the previous step.
5. Click **OK** to save the settings.
6. Click **Start Streaming** to start streaming.

![005](https://github.com/Tencent-RTC/obs-trtc/assets/2777660/612ad0f9-9927-4b48-82da-8ac388dece80)

At this point, the stream is being broadcasted to the TRTC service.

## Step 4: Play the stream

Open the previous webpage which generated the WHIP Bearer Token, go to the `WHEP Player` section, 
and click **Play Stream** to play the stream via WHEP.

![006](https://github.com/Tencent-RTC/obs-trtc/assets/2777660/2e727bce-6d2b-47c8-b214-6fc320b1291a)

Another option is go to the `TRTC Room` section, and click **Join Room** to access the TRTC room 
and watch the stream via TRTC, or you can utilize the TRTC mobile SDK to join the room and view 
the stream.

![007](https://github.com/Tencent-RTC/obs-trtc/assets/2777660/0a18bd36-e1f3-4070-bc89-95218785cb95)

Since both WHIP and WHEP are standard protocols, you can utilize any client that supports 
them to play the stream.

## Conclusion

We explored the use of elastic cloud services for a more stable streaming platform and the 
steps to set up a TRTC application with OBS WHIP support. These technologies enable efficient 
containerization and seamless real-time streaming experiences for various use cases.

If you require assistance or encounter any difficulties, please feel free to reach out
to us via [Discord](https://discord.gg/vDHty6ddrZ).

In the future, we will further investigate specific use cases and technologies associated
with OBS WHIP and TRTC integration.

