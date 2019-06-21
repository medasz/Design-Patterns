package main

type MediaPlayer interface {
	play(audioType string, filename string)
}

//---------//---------//---------//---------//---------
type AdvancedMediaPlayer interface {
	playVlc(filename string)
	playMp4(filename string)
}

type VlcPlayer struct {
}

func (v VlcPlayer) playVlc(filename string) {
	println("Playing vlc file.Name:" + filename)
}
func (v VlcPlayer) playMp4(filename string) {
}

type Mp4Player struct {
}

func (m Mp4Player) playVlc(filename string) {
}
func (m Mp4Player) playMp4(filename string) {
	println("Playing mp4 file.name:" + filename)
}

//---------//---------//---------//---------//---------

type MediaAdapter struct {
	advancedMusicPlayer AdvancedMediaPlayer
}

func (m *MediaAdapter) NewMediaAdapter(audioType string) {
	if audioType == "vlc" {
		m.advancedMusicPlayer = new(VlcPlayer)
	} else if audioType == "mp4" {
		m.advancedMusicPlayer = new(Mp4Player)
	}
}
func (m MediaAdapter) play(audioType string, fileName string) {
	if audioType == "vlc" {
		m.advancedMusicPlayer.playVlc(fileName)
	} else if audioType == "mp4" {
		m.advancedMusicPlayer.playMp4(fileName)
	}
}

type AudioPlayer struct {
	mediaAdapter MediaAdapter
}

func (a AudioPlayer) play(audioType string, fileName string) {
	if audioType == "mp3" {
		println("Playing mp3 file.Name:" + fileName)
	} else if audioType == "vlc" || audioType == "mp4" {
		a.mediaAdapter = MediaAdapter{}
		a.mediaAdapter.NewMediaAdapter(audioType)
		a.mediaAdapter.play(audioType, fileName)
	} else {
		println("Invalid media" + audioType + "format not supported")
	}
}

func main() {
	var a MediaPlayer
	a = AudioPlayer{}
	a.play("mp3", "beyond the horizon.mp3")
	a.play("mp4", "alone.mp4")
	a.play("vlc", "far far away.vlc")
	a.play("avi", "mid me.avi")
}
