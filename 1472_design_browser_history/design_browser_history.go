package main

type Entry struct {
	url  string
	prev *Entry
	next *Entry
}

type BrowserHistory struct {
	home    *Entry
	current *Entry
}

func Constructor(homepage string) BrowserHistory {
	firstNode := &Entry{url: homepage}
	return BrowserHistory{
		home:    firstNode,
		current: firstNode,
	}
}

func (bh *BrowserHistory) Visit(url string) {
	bh.current.next = nil
	newEntry := &Entry{
		url:  url,
		prev: bh.current,
	}
	bh.current.next = newEntry
	bh.current = bh.current.next
}

func (bh *BrowserHistory) Back(steps int) string {
	for i := 0; i < steps && bh.current.prev != nil; i++ {
		bh.current = bh.current.prev
	}
	return bh.current.url
}

func (bh *BrowserHistory) Forward(steps int) string {
	for i := 0; i < steps && bh.current.next != nil; i++ {
		bh.current = bh.current.next
	}
	return bh.current.url
}
