package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type track struct {
	artist string
	age int
	musicTitle string
	year int
	musicLength time.Duration
}

type sortByArtist []*track

func (t sortByArtist) Len() int {return len(t)} //单条目的长度 （用于重新分配条目Buffer?）
func (t sortByArtist) Less(i, j int) bool { return t[i].artist <  t[j].artist} //定义排序规则
func (t sortByArtist) Swap(i, j int)      { t[i], t[j] =t[j], t[i] }

type customSort struct {
	t    []*track
	less func(x, y *track) bool
}
func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)    { x.t[i], x.t[j] = x.t[j], x.t[i] }

var musicList = []*track {
	{"Mary", 28, "kill this love", 2024,length("5m28s")},
	{"Mary", 28, "kill this love2", 2022,length("4m28s")},
	{"Bob", 18, "Umbrella", 2012,length("2m19s")},
	{"Alice", 20, "Hello", 2016,length("4m10s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Artist", "Age", "Music", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.artist, t.age, t.musicTitle, t.year, t.musicLength)
	}
	tw.Flush() // calculate column widths and print table
}

func main() {
	sort.Sort(sortByArtist(musicList))
	printTracks(musicList)
	//如果需要逆向排序 不需要重新定义颠倒less方法的新类型 sort包中支持逆向
	sort.Sort(sort.Reverse(sortByArtist(musicList)))
	printTracks(musicList)
	/*
	sort.Reverse的实现
	Reverse Func返回了reverse类型(reverse不公开 需要通过调用Reverse函数进行转换)
	基于reverse类型重新实现了less接口 将索引反转
	而Len Swap接口隐式的由内嵌的sort.interface提供
	 */

	/*
	当我们需要按照year排序  需要再
	type sortByYear []*track
	func (x sortByYear)  Len() int { return len(x.t) }
	.....
	.....
	实际无论按照什么排序 ，Len和Swap方法的实现都是一样的，因此再我们需要按照不同种类排序的时候，不停的重新定义
	新类型，重新定义基于新类型的方法就显得很累赘。
	可以采用如下方式：
	每次仅需要重新定义less函数 按照想要的标准去的排序
	 */
	sort.Sort(customSort{musicList, func(x, y *track) bool {
		//三个级别排序
		if x.artist != y.artist {
			return x.artist < y.artist
		}

		if x.year != y.year {
			return x.year < y.year
		}

		if x.musicLength != y.musicLength {
			return x.musicLength < y.musicLength
		}
		return false
	}})
	printTracks(musicList)

	values := []int{3,1,4,2}
	fmt.Println(sort.IntsAreSorted(values))
	sort.Ints(values)
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
}
