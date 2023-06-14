# OBS-TRTC

[![](https://img.shields.io/twitter/follow/TencentRTC?style=social)](https://twitter.com/TencentRTC)
[![](https://img.shields.io/badge/TencentRTC-YouTube-red)](https://www.youtube.com/@TencentRTC)
[![](https://badgen.net/discord/members/vDHty6ddrZ)](https://discord.gg/vDHty6ddrZ)

Demo for using OBS WHIP to publish a stream to the TRTC(Tencent Real-Time Communication) service.

## Overview

OBS includes WHIP support, which allows you to do many interesting things by combining 
the powers of both OBS and WHIP.

WHIP is a standard protocol that lets you use HTML5 and different clients to publish 
and play live streams. Plus, you can use open-source tools to build your own live 
streaming platform.

You can also use TRTC (Tencent Real-Time Communication) cloud services with OBS WHIP 
support for a streaming platform. This is a great option if you don't want to build 
your own platform or need a more reliable and stable platform with dedicated support.

Additionally, TRTC (Tencent Real-Time Communication) provides a free trial that includes 
a specific amount of streaming time, making it super easy for you to try out.

If you need help or run into any problems, don't hesitate to contact us on 
[Discord](https://discord.gg/vDHty6ddrZ).

## Prerequisites

Before you move forward, double-check that you've got these necessary items ready:

- OBS with WHIP support, please download from [OBS](https://obsproject.com/)
- TRTC(Tencent Real-Time Communication) account, please register at [here](https://www.tencentcloud.com/account/login?s_url=https%3A%2F%2Fconsole.tencentcloud.com%2Ftrtc&from_portal=click_trtc_home&ext1=Get_Started_for_Free)

> Note: Currently, OBS WHIP has been merged into the master branch but has not been released yet. 
> You can login GitHub and download it from [here](https://github.com/obsproject/obs-studio/actions/runs/5227109208?pr=7926).

Next, you need to create a TRTC application and generate a Bearer Token for WHIP.

## Step 1: Create a TRTC application

Please follow the steps below to create a TRTC application:

1. Log in to the [TRTC console](https://console.cloud.tencent.com/trtc) and click **Application Management** on the left sidebar.
2. Click **Create Application** and enter the application name.
3. Click **Create** to create the application.

![001](https://github.com/Tencent-RTC/obs-trtc/assets/2777660/dce31494-ac4f-4844-b437-de3d244af678)

After the application is created, you can view the SDKAppID and SDKSecretKey on the application details page.

![003](https://github.com/Tencent-RTC/obs-trtc/assets/2777660/f54ad443-645a-4089-8dc7-5a34af48a335)

* SDKAppID: `2000xxxx`
* SDKSecretKey: `yyyyyy`

Following that, you must generate a Bearer Token for WHIP, which will be 
utilized in OBS.

## Step 2: Create a Bearer Token for WHIP

You can directly visit https://tencent-rtc.github.io/obs-trtc/ to create a WHIP Bearer Token. 
Ensure that use the appid with your own `SDKAppID` and secret with your own `SDKSecretKey`, 
then click the `Generate Bearer Token` button.

![0031](https://github.com/Tencent-RTC/obs-trtc/assets/2777660/a6df5559-570e-4348-8492-60ffd7aaa2de)

> Note: You can also access the url `https://tencent-rtc.github.io/obs-trtc/?appid=2000xxx&secret=yyyyyy` to setup the parameters.

Next, use the generated WHIP Bearer Token to configure OBS.

## Step 3: Configure OBS

In the `OBS WHIP` section, you will find the generated WHIP Bearer Token for configuring OBS.

![004](https://github.com/Tencent-RTC/obs-trtc/assets/2777660/29f1ac74-cb3a-4c6e-a3e4-e6c92ea2aab0)

* Service: `WHIP`
* Server: `https://signaling.rtc.tencentcloud.com/v2/pub/2000xxxx/xxxxxxx?SessionID=xxxxxxxxx`
* Bearer Token: `xxxxxx-yyyyyy-zzzzzz`

Please follow the steps below to configure OBS:

1. Open OBS and click **Settings**.
2. Click **Stream** on the left sidebar.
3. Select `WHIP` for **Service**.
4. Enter the `Server` and `Bearer Token` generated in the previous step.
5. Click **OK** to save the settings.
6. Click **Start Streaming** to start streaming.

![005](https://github.com/Tencent-RTC/obs-trtc/assets/2777660/612ad0f9-9927-4b48-82da-8ac388dece80)

At this point, the stream is streaming to the TRTC service.

## Step 4: Play the stream

Open the previous webpage at `tencent-rtc.github.io` which generated the WHIP Bearer Token, go to the `WHEP Player` section, 
and click **Play Stream** to play the stream via WHEP.

![006](https://github.com/Tencent-RTC/obs-trtc/assets/2777660/2e727bce-6d2b-47c8-b214-6fc320b1291a)

Another option is go to the `TRTC Room` section, and click **Join Room** to access the TRTC room 
and watch the stream via TRTC, or you can utilize the TRTC mobile SDK to join the room and view 
the stream.

![007](https://github.com/Tencent-RTC/obs-trtc/assets/2777660/0a18bd36-e1f3-4070-bc89-95218785cb95)

Since both WHIP and WHEP are standard protocols, you can utilize any client that supports 
them to play the stream.

## Conclusion

We looked into using flexible TRTC (Tencent Real-Time Communication) cloud services to make a stronger 
streaming platform and the steps needed to create a TRTC app with OBS WHIP help. These tools make it 
easier to organize and provide smooth live streaming experiences for different situations.

If you require assistance or encounter any difficulties, please feel free to reach out
to us via [Discord](https://discord.gg/vDHty6ddrZ).

In the future, we will further investigate specific use cases and technologies associated
with OBS WHIP and TRTC integration.

## (Optional) Annexb A: Deploy the WHIP Bearer Token Server

We deploy this tool on the server `https://tencent-rtc.github.io/obs-trtc/` for your convenience.

Alternatively, you can operate your own server to produce the token by utilizing
the following command line, then navigate to the webpage `http://localhost:9000/`
to generate a WHIP Bearer Token.

```bash
git clone https://github.com/Tencent-RTC/obs-trtc.git
cd obs-trtc
open index.html
```

The other steps are the same as the previous section.

## (Optional) Annexb B: Configure OBS in Real-Time Mode

If you want to configure OBS in real-time mode, please follow the steps below:

1. Open OBS and click **Settings**.
2. Click **Output** on the left sidebar.
3. Select `Advanced` for **Output Mode**.
4. Set **Keyframe Interval** to `1 s`.
5. Select `veryfast` for **CPU Usage Preset**.
6. Select `baseline` for **Profile**.
7. Select `zerolatency` for **Tune**.
8. Click **OK** to save the settings.

![009](https://github.com/Tencent-RTC/obs-trtc/assets/2777660/76c36f45-f164-4a78-b3fc-29005cacf6e7)

The other steps are the same as the previous section.
