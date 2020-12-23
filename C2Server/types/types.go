package types

type TestStruct struct {
	Wadu string `json:wadu`
	Waduagain []int `json:waduagain`
}

type LoggerMain struct {
	TimeStamp string `json:time_stamp`
	NanosecondsSinceEpoch int64 `json:nanoseconds_since_epoch`
	Msg string `json:msg`
	Author string `json:author`
	Fields map[string]interface{} `json:fields`
	File string `json:file`
	Function string `json:function`
	Line int `json:line`
	Level string `json:level`
}

type TimeObject struct {
	Seconds int64 `json:seconds`
	TimeStamp string `json:time_stamp`
}
type TimesObject struct {
	Duration TimeObject `json:duration`
	CurrentPosition TimeObject `json:duration`
	Remaining TimeObject `json:duration`
}
type StatsObject struct {
	Skipped bool `json:skipped`
	NumberOfTimesSkipped int `json:number_of_times_skipped`
	Watched bool `json:watched`
	NumberOfTimesWatched int `json:number_of_times_watched`
	Completed bool `json:completed`
	NumberOfTimesCompleted int `json:number_of_times_completed`
}
type NowPlayingMeta struct {
	Title string `json:title`
	Artist string `json:artist`
	LocalFilePath string `json:local_file_path`
	LocalFilePathB64 string `json:local_file_path_b64`
	URL string `json:url`
	ShowIndex string `json:show_index`
	EpisodeIndex string `json:episode_index`
	Times TimesObject `json:times`
	Stats StatsObject `json:stats`
}

type StateMetaData struct {
	Name string `json:name`
	GenericType string `json:generic_type`
	RestartOnFail bool `json:restart_on_fail`
	NowPlaying NowPlayingMeta `json:now_playing`
}

type VLCStatus struct {
	Input string `json:input`
	AudioVolume int `json:audio_volume`
	State string `json:state`
}

type VLCSInfo string

type VLCCommonStatus struct {
	Input string `json:input`
	AudioVolume int `json:audio_volume`
	State string `json:state`
	Info string `json:string`
	Seconds int `json:seconds`
	Length int `json:length`
}

type SpotifyStatus struct {
	Shuffle string
	MaximumRate string
	MinimumRate string
	Rate string
	Volume float64
	Position int64
	LoopStatus string
	Playback string
	Metadata struct {
		TrackID string
		Artist []string
		Title string
		Album string
		TrackNumber int32
		Rating int
		Status string
		Url string
		ArtUrl string
		ArtFile string
	}
}