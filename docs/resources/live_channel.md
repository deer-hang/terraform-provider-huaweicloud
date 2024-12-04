---
subcategory: "Live"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_live_channel"
description: |-
  Manages a Live channel resource within HuaweiCloud.
---

# huaweicloud_live_channel

Manages a Live channel resource within HuaweiCloud.

## Example Usage

### Create a Live channel with RTMP_PUSH protocol

```hcl
variable "domain_name" {}
variable "template_id" {}
variable "hls_package_url" {}
variable "hls_package_encryption_url" {}
variable "hls_package_encryption_urn" {}
variable "mss_package_url" {}

resource "huaweicloud_live_channel" "test" {
  app_name    = "live"
  domain_name = var.domain_name
  name        = "test-name"
  state       = "ON"

  encoder_settings {
    template_id = var.template_id
  }

  endpoints {
    hls_package {
      hls_version              = "v3"
      playlist_window_seconds  = 24
      segment_duration_seconds = 4
      url                      = var.hls_package_url

      encryption {
        encryption_method             = "SAMPLE-AES"
        key_rotation_interval_seconds = 0
        level                         = "profile"
        request_mode                  = "functiongraph_proxy"
        resource_id                   = "test-resource-id"
        speke_version                 = "1.0"
        system_ids                    = [
          "FairPlay",
        ]
        url = var.hls_package_encryption_url
        urn = var.hls_package_encryption_urn
      }

      request_args {
        record {
          end_time   = "end"
          format     = "timestamp"
          start_time = "begin"
          unit       = "second"
        }

        timeshift {
          back_time = "delay"
          unit      = "second"
        }
      }
    }

    mss_package {
      playlist_window_seconds  = 42
      segment_duration_seconds = 4
      url                      = var.mss_package_url

      encryption {
        encryption_method             = "SAMPLE-AES"
        key_rotation_interval_seconds = 0
        level                         = "content"
        request_mode                  = "direct_http"
        resource_id                   = "test-resource-id"
        speke_version                 = "1.0"
        system_ids                    = [
          "PlayReady",
        ]
        url = "https://test-url.cd"
      }

      request_args {
        record {
          end_time   = "end"
          format     = "timestamp"
          start_time = "begin"
          unit       = "second"
        }

        timeshift {
          back_time = "delay"
          unit      = "second"
        }
      }
    }
  }

  input {
    input_protocol = "RTMP_PUSH"

    sources {
      bitrate = 100
    }
  }

  record_settings {
    rollingbuffer_duration = 20
  }
}
```

### Create a Live channel with FLV_PULL protocol

```hcl
variable "domain_name" {}
variable "template_id" {}
variable "dash_package_url" {}
variable "hls_package_url" {}

resource "huaweicloud_live_channel" "test" {
  app_name    = "live"
  domain_name = var.domain_name
  name        = "test-name"
  state       = "ON"

  encoder_settings {
    template_id = var.template_id
  }

  endpoints {
    dash_package {
      playlist_window_seconds  = 24
      segment_duration_seconds = 4
      url                      = var.dash_package_url

      request_args {
        record {
          end_time   = "end"
          format     = "timestamp"
          start_time = "begin"
          unit       = "second"
        }

        timeshift {
          back_time = "delay"
          unit      = "second"
        }
      }
    }

    hls_package {
      hls_version              = "v3"
      playlist_window_seconds  = 8
      segment_duration_seconds = 4
      url                      = var.hls_package_url

      encryption {
        encryption_method             = "SAMPLE-AES"
        key_rotation_interval_seconds = 0
        level                         = "content"
        request_mode                  = "direct_http"
        resource_id                   = "test-resource-id"
        speke_version                 = "1.0"
        system_ids = [
          "FairPlay",
        ]
        url = "http://xxx.sp"

        http_headers {
          key   = "aaa"
          value = "sss"
        }

        http_headers {
          key   = "www"
          value = "qqq"
        }
      }

      request_args {
        record {
          end_time   = "end"
          format     = "timestamp"
          start_time = "begin"
          unit       = "second"
        }

        timeshift {
          back_time = "delay"
          unit      = "second"
        }
      }
    }
  }

  input {
    input_protocol = "FLV_PULL"

    failover_conditions {
      input_loss_threshold_msec = 4000
      input_preference          = "EQUAL"
    }

    secondary_sources {
      bitrate = 100
      url     = "https://hgf.vv"
    }

    sources {
      bitrate = 100
      url     = "https://qwe.cc"
    }
  }

  record_settings {
    rollingbuffer_duration = 12
  }
}
```

### Create a Live channel with HLS_PULL protocol

```hcl
variable "domain_name" {}
variable "template_id" {}
variable "hls_package_url" {}
variable "hls_package_encryption_url" {}
variable "hls_package_encryption_urn" {}
variable "mss_package_url" {}

resource "huaweicloud_live_channel" "test" {
  app_name    = "live"
  domain_name = var.domain_name
  name        = "test-name"
  state       = "ON"

  encoder_settings {
    template_id = var.template_id
  }

  endpoints {
    hls_package {
      hls_version              = "v3"
      playlist_window_seconds  = 4
      segment_duration_seconds = 2
      url                      = var.hls_package_url

      encryption {
        encryption_method             = "SAMPLE-AES"
        key_rotation_interval_seconds = 0
        level                         = "content"
        request_mode                  = "functiongraph_proxy"
        resource_id                   = "test-resource-id"
        speke_version                 = "1.0"
        system_ids = [
          "FairPlay",
        ]
        url = var.hls_package_encryption_url
        urn = var.hls_package_encryption_urn
      }

      request_args {
        record {
          end_time   = "end"
          format     = "timestamp"
          start_time = "begin"
          unit       = "second"
        }

        timeshift {
          back_time = "delay"
          unit      = "second"
        }
      }
    }

    mss_package {
      playlist_window_seconds  = 8
      segment_duration_seconds = 2
      url                      = var.mss_package_url

      encryption {
        encryption_method             = "SAMPLE-AES"
        key_rotation_interval_seconds = 0
        level                         = "content"
        request_mode                  = "direct_http"
        resource_id                   = "dfge"
        speke_version                 = "1.0"
        system_ids = [
          "PlayReady",
        ]
        url = "https://ssc.cd"

        http_headers {
          key   = "aa"
          value = "ss"
        }

        http_headers {
          key   = "gg"
          value = "ff"
        }
      }

      request_args {
        record {
          end_time   = "end"
          format     = "timestamp"
          start_time = "begin"
          unit       = "second"
        }

        timeshift {
          back_time = "delay"
          unit      = "second"
        }
      }
    }
  }

  input {
    input_protocol      = "HLS_PULL"
    max_bandwidth_limit = 200

    failover_conditions {
      input_loss_threshold_msec = 2000
      input_preference          = "PRIMARY"
    }

    secondary_sources {
      bitrate = 100
      url     = "https://qqwe.dd"
    }

    sources {
      bitrate = 100
      url     = "https://ssa.qw"
    }
  }

  record_settings {
    rollingbuffer_duration = 3
  }
}
```

### Create a Live channel with SRT_PUSH protocol

```hcl
variable "domain_name" {}
variable "template_id" {}
variable "hls_package_url" {}
variable "mss_package_url" {}

resource "huaweicloud_live_channel" "test" {
  app_name    = "live"
  domain_name = var.domain_name
  name        = "test-name"
  state       = "ON"

  encoder_settings {
    template_id = var.template_id
  }
 
  endpoints {
    hls_package {
      hls_version              = "v3"
      playlist_window_seconds  = 4
      segment_duration_seconds = 2
      url                      = var.hls_package_url

      encryption {
        encryption_method             = "SAMPLE-AES"
        key_rotation_interval_seconds = 0
        level                         = "content"
        request_mode                  = "direct_http"
        resource_id                   = "test-resource-id"
        speke_version                 = "1.0"
        system_ids = [
          "FairPlay",
        ]
        url = "https://qqq.co"

        http_headers {
          key   = "aa"
          value = "sss"
        }

        http_headers {
          key   = "dd"
          value = "sss"
        }
      }

      request_args {
        record {
          end_time   = "end"
          format     = "timestamp"
          start_time = "begin"
          unit       = "second"
        }

        timeshift {
          back_time = "delay"
          unit      = "second"
        }
      }
    }

    mss_package {
      playlist_window_seconds  = 8
      segment_duration_seconds = 2
      url                      = var.mss_package_url

      request_args {
        record {
          end_time   = "end"
          format     = "timestamp"
          start_time = "begin"
          unit       = "second"
        }

        timeshift {
          back_time = "delay"
          unit      = "second"
        }
      }
    }
  }

  input {
    input_protocol = "SRT_PUSH"
    ip_whitelist   = "192.168.0.1/16,192.168.1.1/16,192.168.2.1/16"

    audio_selectors {
      name = "test-audio-selectors1"

      selector_settings {
        audio_pid_selection {
          pid = 2
        }
      }
    }

    audio_selectors {
      name = "test-audio-selectors2"

      selector_settings {
        audio_language_selection {
          language_code             = "ch"
          language_selection_policy = "LOOSE"
        }
      }
    }

    audio_selectors {
      name = "test-audio-selectors3"

      selector_settings {
        audio_pid_selection {
          pid = 0
        }
      }
    }

    sources {
      bitrate = 100
    }
  }

  record_settings {
    rollingbuffer_duration = 2
  }
}
```

### Create a Live channel with SRT_PULL protocol

```hcl
variable "domain_name" {}
variable "template_id" {}
variable "hls_package_url" {}
variable "mss_package_url" {}

resource "huaweicloud_live_channel" "test" {
  app_name    = "live"
  domain_name = var.domain_name
  name        = "test-name"
  state       = "ON"

  encoder_settings {
    template_id = var.template_id
  }

  endpoints {
    hls_package {
      hls_version              = "v3"
      playlist_window_seconds  = 16
      segment_duration_seconds = 4
      url                      = var.hls_package_url

      encryption {
        encryption_method = "SAMPLE-AES"
        level             = "content"
        request_mode      = "direct_http"
        resource_id       = "test-resource-id"
        speke_version     = "1.0"
        system_ids = [
          "FairPlay",
        ]
        url = "https://sss.cc"

        http_headers {
          key   = "aa"
          value = "ss"
        }

        http_headers {
          key   = "ff"
          value = "dd"
        }
      }

      request_args {
        record {
          end_time   = "end"
          format     = "timestamp"
          start_time = "begin"
          unit       = "second"
        }

        timeshift {
          back_time = "delay"
          unit      = "second"
        }
      }
    }

    mss_package {
      playlist_window_seconds  = 24
      segment_duration_seconds = 4
      url                      = var.mss_package_url

      request_args {
        record {
          end_time   = "end"
          format     = "timestamp"
          start_time = "begin"
          unit       = "second"
        }

        timeshift {
          back_time = "delay"
          unit      = "second"
        }
      }
    }
  }

  input {
    input_protocol = "SRT_PULL"

    audio_selectors {
      name = "test-audio-selectors1"

      selector_settings {
        audio_language_selection {
          language_code             = "dfg"
          language_selection_policy = "LOOSE"
        }
      }
    }

    audio_selectors {
      name = "test-audio-selectors2"

      selector_settings {
        audio_pid_selection {
          pid = 13
        }
      }
    }

    audio_selectors {
      name = "test-audio-selectors3"

      selector_settings {
        audio_pid_selection {
          pid = 0
        }
      }
    }

    failover_conditions {
      input_loss_threshold_msec = 2000
      input_preference          = "EQUAL"
    }

    secondary_sources {
      bitrate   = 100
      latency   = 1000
      stream_id = "vcbeer"
      url       = "srt://192.168.1.215:9001"
    }

    sources {
      bitrate   = 100
      latency   = 2000
      stream_id = "dfawerw"
      url       = "srt://192.168.1.216:9001"
    }
  }

  record_settings {
    rollingbuffer_duration = 4
  }
}
```

### Create a Live channel with custom channel ID

```hcl
variable "domain_name" {}
variable "channel_id" {}
variable "template_id" {}
variable "hls_package_url" {}

resource "huaweicloud_live_channel" "test" {
  app_name    = "live"
  domain_name = var.domain_name
  state       = "OFF"
  channel_id  = var.channel_id

  encoder_settings {
    template_id = var.template_id
  }

  endpoints {
    hls_package {
      hls_version              = "v3"
      playlist_window_seconds  = 40
      segment_duration_seconds = 4
      url                      = var.hls_package_url

      request_args {
        record {
          end_time   = "end"
          format     = "timestamp"
          start_time = "begin"
          unit       = "second"
        }

        timeshift {
          back_time = "delay"
          unit      = "second"
        }
      }
    }
  }

  input {
    input_protocol = "RTMP_PUSH"

    sources {
      bitrate = 100
    }
  }

  record_settings {
    rollingbuffer_duration = 0
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the resource.
  If omitted, the provider-level region will be used. Changing this parameter will create a new resource.

* `domain_name` - (Required, String) Specifies the channel streaming domain name.

* `app_name` - (Required, String) Specifies the group name or application name. For example: **live**.

* `state` - (Required, String) Specifies the channel status. Valid values are:
  + **ON**: After a channel is delivered, functions such as stream pull, transcoding, and recording are automatically enabled.
  + **OFF**: Only the channel information is saved but the channel is not started.

* `input` - (Required, List) Specifies the channel input information.
The [input](#LiveChannel_Input) structure is documented below.

* `record_settings` - (Required, List) Specifies the configuration for replaying a recording.
The [RecordSettings](#LiveChannel_RecordSettings) structure is documented below.

* `endpoints` - (Required, List) Specifies the channel outflow information.
The [Endpoints](#LiveChannel_Endpoints) structure is documented below.

* `encoder_settings_expand` - (Optional, List) Specifies the audio output configuration.
The [EncoderSettingsExpand](#LiveChannel_EncoderSettingsExpand) structure is documented below.

* `encoder_settings` - (Optional, List) Specifies the transcoding template configuration.
The [EncoderSettings](#LiveChannel_EncoderSettings) structure is documented below.

* `name` - (Optional, String) Specifies the channel name.

* `channel_id` - (Optional, String) Specifies the unique channel ID.

<a name="LiveChannel_Input"></a>
The `input` block supports:

* `input_protocol` - (Required, String) Specifies the channel input protocol. Valid values are:
  + **FLV_PULL**.
  + **RTMP_PUSH**.
  + **HLS_PULL**.
  + **SRT_PULL**.
  + **SRT_PUSH**.

* `sources` - (Optional, List) Specifies the channel main source stream information.
The [Sources](#LiveChannel_Sources) structure is documented below.

* `secondary_sources` - (Optional, List) Specifies the prepared stream array.
The [SecondarySources](#LiveChannel_SecondarySources) structure is documented below.

* `failover_conditions` - (Optional, List) Specifies the configuration of switching between primary and backup audio and video stream URLs.
The [FailoverConditions](#LiveChannel_FailoverConditions) structure is documented below.

* `max_bandwidth_limit` - (Optional, Int) Specifies the maximum bandwidth that needs to be configured when the inbound protocol is **HLS_PULL**.

* `ip_port_mode` - (Optional, Bool) Specifies the IP port mode.

* `ip_whitelist` - (Optional, String) Specifies the IP whitelist when protocol is **SRT_PUSH**.

* `scte35_source` - (Optional, String) Specifies the advertisement scte35 signal source.

* `ad_triggers` - (Optional, List) Specifies the ad trigger configuration.

* `audio_selectors` - (Optional, List) Specifies the audio selector configuration.
The [AudioSelectors](#LiveChannel_AudioSelectors) structure is documented below.

<a name="LiveChannel_InputSources"></a>
The `InputSources` block supports:

* `url` - (Optional, String) Specifies the channel source stream URL, used for external streaming.

* `bitrate` - (Optional, Int) Specifies the bitrate.

* `width` - (Optional, Int) Specifies the resolution corresponds to the width value.

* `height` - (Optional, Int) Specifies the resolution corresponds to the high value.

* `enable_snapshot` - (Optional, Bool) Specifies whether to use this stream to take screenshots.

* `bitrate_for3u8` - (Optional, Bool) Specifies whether to use bitrate to fix the bitrate.

* `passphrase` - (Optional, String) Specifies the encrypted information when the protocol is **SRT_PUSH**.

* `backup_urls` - (Optional, List) Specifies the list of backup stream addresses.

* `stream_id` - (Optional, String) Specifies the stream ID of the stream pull address when the channel type is **SRT_PULL**.

* `latency` - (Optional, Int) Specifies the streaming delay when the channel type is **SRT_PULL**.

<a name="LiveChannel_InputSecondarySources"></a>
The `InputSecondarySources` block supports:

* `url` - (Optional, String) Specifies the channel source stream URL.

* `bitrate` - (Optional, Int) Specifies the bitrate.

* `width` - (Optional, Int) Specifies the resolution corresponds to the width value.

* `height` - (Optional, Int) Specifies the resolution corresponds to the high value.

* `bitrate_for3u8` - (Optional, Bool) Specifies whether to use bitrate to fix the bitrate.

* `passphrase` - (Optional, String) Specifies the encrypted information when the protocol is **SRT_PUSH**.

* `backup_urls` - (Optional, List) Specifies the list of backup stream addresses.

* `stream_id` - (Optional, String) Specifies the stream ID of the stream pull address when the channel type is **SRT_PULL**.

* `latency` - (Optional, Int) Specifies the streaming delay when the channel type is **SRT_PULL**.

<a name="LiveChannel_InputFailoverConditions"></a>
The `InputFailoverConditions` block supports:

* `input_loss_threshold_msec` - (Optional, Int) Specifies the duration threshold of inflow stop.

* `input_preference` - (Optional, String) Specifies the input preference type.

<a name="LiveChannel_InputAudioSelectors"></a>
The `InputAudioSelectors` block supports:

* `name` - (Required, String) Specifies the name of the audio selector.

* `selector_settings` - (Optional, List) Specifies the audio selector configuration.
The [SelectorSettings](#LiveChannel_SelectorSettings) structure is documented below.

<a name="LiveChannel_AudioSelectorsSelectorSettings"></a>
The `AudioSelectorsSelectorSettings` block supports:

* `audio_language_selection` - (Optional, List) Specifies the language selector configuration.
The [AudioLanguageSelection](#LiveChannel_AudioLanguageSelection) structure is documented below.

* `audio_pid_selection` - (Optional, List) Specifies the PID selector configuration.
The [AudioPidSelection](#LiveChannel_AudioPidSelection) structure is documented below.

* `audio_hls_selection` - (Optional, List) Specifies the HLS selector configuration.
The [AudioHlsSelection](#LiveChannel_AudioHlsSelection) structure is documented below.

<a name="LiveChannel_SelectorSettingsAudioLanguageSelection"></a>
The `SelectorSettingsAudioLanguageSelection` block supports:

* `language_code` - (Required, String) Specifies the language abbreviation.

* `language_selection_policy` - (Optional, String) Specifies the language output strategy.

<a name="LiveChannel_SelectorSettingsAudioPidSelection"></a>
The `SelectorSettingsAudioPidSelection` block supports:

* `pid` - (Required, Int) Specifies the value of PID.

<a name="LiveChannel_SelectorSettingsAudioHlsSelection"></a>
The `SelectorSettingsAudioHlsSelection` block supports:

* `name` - (Required, String) Specifies the HLS audio selector name.

* `group_id` - (Required, String) Specifies the HLS audio selector gid.

<a name="LiveChannel_RecordSettings"></a>
The `RecordSettings` block supports:

* `rollingbuffer_duration` - (Required, Int) Specifies the maximum playback recording time.

<a name="LiveChannel_Endpoints"></a>
The `Endpoints` block supports:

* `hls_package` - (Optional, List) Specifies the HLS packaging information.
The [HlsPackage](#LiveChannel_HlsPackage) structure is documented below.

* `dash_package` - (Optional, List) Specifies the DASH packaging information.
The [DashPackage](#LiveChannel_DashPackage) structure is documented below.

* `mss_package` - (Optional, List) Specifies the MSS packaging information.
The [MssPackage](#LiveChannel_MssPackage) structure is documented below.

<a name="LiveChannel_EndpointsHlsPackage"></a>
The `EndpointsHlsPackage` block supports:

* `url` - (Required, String) Specifies the customer-defined streaming address.

* `stream_selection` - (Optional, List) Specifies the stream selection.
The [StreamSelection](#LiveChannel_StreamSelection) structure is documented below.

* `hls_version` - (Optional, String) Specifies the HLS version number.

* `segment_duration_seconds` - (Required, Int) Specifies the duration of the channel output segment.

* `playlist_window_seconds` - (Optional, Int) Specifies the window length of the channel live broadcast return shard.

* `encryption` - (Optional, List) Specifies the encrypted information.
The [Encryption](#LiveChannel_Encryption) structure is documented below.

* `request_args` - (Optional, List) Specifies the play related configuration.
The [RequestArgs](#LiveChannel_RequestArgs) structure is documented below.

* `ad_marker` - (Optional, List) Specifies the advertising marker.

<a name="LiveChannel_HlsPackageStreamSelection"></a>
The `HlsPackageStreamSelection` block supports:

* `key` - (Optional, String) Specifies the key used for bitrate filtering in streaming URLs.

* `max_bandwidth` - (Optional, Int) Specifies the maximum code rate.

* `min_bandwidth` - (Optional, Int) Specifies the minimum code rate.

<a name="LiveChannel_HlsPackageEncryption"></a>
The `HlsPackageEncryption` block supports:

* `key_rotation_interval_seconds` - (Optional, Int) Specifies the key rotation interval seconds.

* `encryption_method` - (Optional, String) Specifies the encryption method.

* `level` - (Optional, String) Specifies the level.

* `resource_id` - (Required, String) Specifies the customer-generated DRM content ID.

* `system_ids` - (Required, List) Specifies the system ID enumeration values.

* `url` - (Required, String) Specifies the DRM address of the key.

* `speke_version` - (Required, String) Specifies the DRM spec version number.

* `request_mode` - (Required, String) Specifies the request mode.

* `http_headers` - (Optional, List) Specifies the authentication information that needs to be added to the DRM request header.
The [HttpHeader](#LiveChannel_HttpHeader) structure is documented below.

* `urn` - (Optional, String) Specifies the URN of the function graph.

<a name="LiveChannel_EncryptionHttpHeader"></a>
The `EncryptionHttpHeader` block supports:

* `key` - (Required, String) Specifies the key field name in the request header.

* `value` - (Required, String) Specifies the value corresponding to the key in the request header.

<a name="LiveChannel_HlsPackageRequestArgs"></a>
The `HlsPackageRequestArgs` block supports:

* `record` - (Optional, List) Specifies the recording and playback related configuration.
The [Record](#LiveChannel_Record) structure is documented below.

* `timeshift` - (Optional, List) Specifies the time-shift playback configuration.
The [TimeShift](#LiveChannel_TimeShift) structure is documented below.

* `live` - (Optional, List) Specifies the live broadcast configuration.
The [Live](#LiveChannel_Live) structure is documented below.

<a name="LiveChannel_RequestArgsRecord"></a>
The `RequestArgsRecord` block supports:

* `start_time` - (Optional, String) Specifies the start time.

* `end_time` - (Optional, String) Specifies the end time.

* `format` - (Optional, String) Specifies the format.

* `unit` - (Optional, String) Specifies the unit.

<a name="LiveChannel_RequestArgsTimeShift"></a>
The `RequestArgsTimeShift` block supports:

* `back_time` - (Optional, String) Specifies the time shift duration field name.

* `unit` - (Optional, String) Specifies the unit.

<a name="LiveChannel_RequestArgsLive"></a>
The `RequestArgsLive` block supports:

* `delay` - (Optional, String) Specifies the delay field.

* `unit` - (Optional, String) Specifies the unit.

<a name="LiveChannel_EndpointsDashPackage"></a>
The `EndpointsDashPackage` block supports:

* `url` - (Required, String) Specifies the customer-defined streaming address.

* `stream_selection` - (Optional, List) Specifies the stream selection.
The [StreamSelection](#LiveChannel_StreamSelection) structure is documented below.

* `segment_duration_seconds` - (Required, Int) Specifies the duration of the channel output segment.

* `playlist_window_seconds` - (Optional, Int) Specifies the window length of the channel live broadcast return shard.

* `encryption` - (Optional, List) Specifies the encrypted information.
The [Encryption](#LiveChannel_Encryption) structure is documented below.

* `request_args` - (Optional, List) Specifies the play related configuration.
The [RequestArgs](#LiveChannel_RequestArgs) structure is documented below.

* `ad_marker` - (Optional, String) Specifies the advertising marker.

<a name="LiveChannel_DashPackageStreamSelection"></a>
The `DashPackageStreamSelection` block supports:

* `key` - (Optional, String) Specifies the key used for bitrate filtering in streaming URLs.

* `max_bandwidth` - (Optional, Int) Specifies the maximum code rate.

* `min_bandwidth` - (Optional, Int) Specifies the minimum code rate.

<a name="LiveChannel_DashPackageEncryption"></a>
The `DashPackageEncryption` block supports:

* `key_rotation_interval_seconds` - (Optional, Int) Specifies the key rotation interval seconds.

* `encryption_method` - (Optional, String) Specifies the encryption method.

* `level` - (Optional, String) Specifies the level.

* `resource_id` - (Required, String) Specifies the customer-generated DRM content ID.

* `system_ids` - (Required, List) Specifies the system ID enumeration values.

* `url` - (Required, String) Specifies the DRM address of the key.

* `speke_version` - (Required, String) Specifies the DRM spec version number.

* `request_mode` - (Required, String) Specifies the request mode.

* `http_headers` - (Optional, List) Specifies the authentication information that needs to be added to the DRM request header.
The [HttpHeader](#LiveChannel_HttpHeader) structure is documented below.

* `urn` - (Optional, String) Specifies the URN of the function graph.

<a name="LiveChannel_EncryptionHttpHeader"></a>
The `EncryptionHttpHeader` block supports:

* `key` - (Required, String) Specifies the key field name in the request header.

* `value` - (Required, String) Specifies the value corresponding to the key in the request header.

<a name="LiveChannel_DashPackageRequestArgs"></a>
The `DashPackageRequestArgs` block supports:

* `record` - (Optional, List) Specifies the recording and playback related configuration.
The [Record](#LiveChannel_Record) structure is documented below.

* `timeshift` - (Optional, List) Specifies the time-shift playback configuration.
The [TimeShift](#LiveChannel_TimeShift) structure is documented below.

* `live` - (Optional, List) Specifies the live broadcast configuration.
The [Live](#LiveChannel_Live) structure is documented below.

<a name="LiveChannel_RequestArgsRecord"></a>
The `RequestArgsRecord` block supports:

* `start_time` - (Optional, String) Specifies the start time.

* `end_time` - (Optional, String) Specifies the end time.

* `format` - (Optional, String) Specifies the format.

* `unit` - (Optional, String) Specifies the unit.

<a name="LiveChannel_RequestArgsTimeShift"></a>
The `RequestArgsTimeShift` block supports:

* `back_time` - (Optional, String) Specifies the time shift duration field name.

* `unit` - (Optional, String) Specifies the unit.

<a name="LiveChannel_RequestArgsLive"></a>
The `RequestArgsLive` block supports:

* `delay` - (Optional, String) Specifies the delay field.

* `unit` - (Optional, String) Specifies the unit.

<a name="LiveChannel_EndpointsMssPackage"></a>
The `EndpointsMssPackage` block supports:

* `url` - (Required, String) Specifies the customer-defined streaming address.

* `stream_selection` - (Optional, List) Specifies the stream selection.
The [StreamSelection](#LiveChannel_StreamSelection) structure is documented below.

* `segment_duration_seconds` - (Required, Int) Specifies the duration of the channel output segment.

* `playlist_window_seconds` - (Optional, Int) Specifies the window length of the channel live broadcast return shard.

* `encryption` - (Optional, List) Specifies the encrypted information.
The [Encryption](#LiveChannel_Encryption) structure is documented below.

* `delay_segment` - (Optional, Int) Specifies the delayed playback time.

* `request_args` - (Optional, List) Specifies the play related configuration.
The [RequestArgs](#LiveChannel_RequestArgs) structure is documented below.

<a name="LiveChannel_MssPackageStreamSelection"></a>
The `MssPackageStreamSelection` block supports:

* `key` - (Optional, String) Specifies the key used for bitrate filtering in streaming URLs.

* `max_bandwidth` - (Optional, Int) Specifies the maximum code rate.

* `min_bandwidth` - (Optional, Int) Specifies the minimum code rate.

<a name="LiveChannel_MssPackageEncryption"></a>
The `MssPackageEncryption` block supports:

* `key_rotation_interval_seconds` - (Optional, Int) Specifies the key rotation interval seconds.

* `encryption_method` - (Optional, String) Specifies the encryption method.

* `level` - (Optional, String) Specifies the level.

* `resource_id` - (Required, String) Specifies the customer-generated DRM content ID.

* `system_ids` - (Required, List) Specifies the system ID enumeration values.

* `url` - (Required, String) Specifies the DRM address of the key.

* `speke_version` - (Required, String) Specifies the DRM spec version number.

* `request_mode` - (Required, String) Specifies the request mode.

* `http_headers` - (Optional, List) Specifies the authentication information that needs to be added to the DRM request header.
The [HttpHeader](#LiveChannel_HttpHeader) structure is documented below.

* `urn` - (Optional, String) Specifies the URN of the function graph.

<a name="LiveChannel_EncryptionHttpHeader"></a>
The `EncryptionHttpHeader` block supports:

* `key` - (Required, String) Specifies the key field name in the request header.

* `value` - (Required, String) Specifies the value corresponding to the key in the request header.

<a name="LiveChannel_MssPackageRequestArgs"></a>
The `MssPackageRequestArgs` block supports:

* `record` - (Optional, List) Specifies the recording and playback related configuration.
The [Record](#LiveChannel_Record) structure is documented below.

* `timeshift` - (Optional, List) Specifies the time-shift playback configuration.
The [TimeShift](#LiveChannel_TimeShift) structure is documented below.

* `live` - (Optional, List) Specifies the live broadcast configuration.
The [Live](#LiveChannel_Live) structure is documented below.

<a name="LiveChannel_RequestArgsRecord"></a>
The `RequestArgsRecord` block supports:

* `start_time` - (Optional, String) Specifies the start time.

* `end_time` - (Optional, String) Specifies the end time.

* `format` - (Optional, String) Specifies the format.

* `unit` - (Optional, String) Specifies the unit.

<a name="LiveChannel_RequestArgsTimeShift"></a>
The `RequestArgsTimeShift` block supports:

* `back_time` - (Optional, String) Specifies the time shift duration field name.

* `unit` - (Optional, String) Specifies the unit.

<a name="LiveChannel_RequestArgsLive"></a>
The `RequestArgsLive` block supports:

* `delay` - (Optional, String) Specifies the delay field.

* `unit` - (Optional, String) Specifies the unit.

<a name="LiveChannel_EncoderSettingsExpand"></a>
The `EncoderSettingsExpand` block supports:

* `audio_descriptions` - (Optional, List) Specifies the description of the audio output configuration.
The [AudioDescriptions](#LiveChannel_AudioDescriptions) structure is documented below.

<a name="LiveChannel_EncoderSettingsExpandAudioDescriptions"></a>
The `EncoderSettingsExpandAudioDescriptions` block supports:

* `name` - (Required, String) Specifies the name of the audio output configuration.

* `audio_selector_name` - (Required, String) Specifies the audio selector name.

* `language_code_control` - (Optional, String) Specifies the language code control configuration.

* `language_code` - (Optional, String) Specifies the language code.

* `stream_name` - (Optional, String) Specifies the stream name.

<a name="LiveChannel_EncoderSettings"></a>
The `EncoderSettings` block supports:

* `template_id` - (Optional, String) Specifies the transcoding template ID.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID.

## Import

The live channel can be imported using the `id`, e.g.

```bash
$ terraform import huaweicloud_live_channel.test 0ce123456a00f2591fabc00385ff1234
```
