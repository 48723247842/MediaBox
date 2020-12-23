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
	Shuffle string `json:shuffle`
	MaximumRate string `json:maximum_rate`
	MinimumRate string `json:minimum_rate`
	Rate string `json:rate`
	Volume float64 `json:volume`
	Position int64 `json:position`
	LoopStatus string `json:loop_status`
	Playback string `json:playback`
	Metadata struct {
		TrackID string `json:track_id`
		Artist []string `json:artist`
		Title string `json:title`
		Album string `json:album`
		TrackNumber int32 `json:track_number`
		Rating int `json:rating`
		Status string `json:status`
		Url string `json:url`
		ArtUrl string `json:art_url`
		ArtFile string `json:art_file`
	} `json:metadata`
}

type TVState struct {
	Power int `json:power`
	Volume int `json:volume`
	Input string `json:input`
	Mute string `json:mute`
}