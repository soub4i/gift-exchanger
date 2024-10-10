package datastore

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Member struct {
	ID      string   `json:"id,omitempty"`
	Name    string   `json:"name"`
	History []string `json:",omitempty"`
}

type Exchange struct {
	Member    string `json:"member_id"`
	Recipient string `json:"recipient_member_id"`
}

type DataStore struct {
	Members   []Member
	Exchanges map[string]string
	Year      int
	mu        sync.Mutex
}

var ds = &DataStore{}

func (d *DataStore) AddMember(n Member) *Member {
	d.mu.Lock()
	defer d.mu.Unlock()
	n.ID = strconv.Itoa(len(d.Members) + 1)
	d.Members = append(d.Members, n)
	return &n
}

func (d *DataStore) DeleteMember(id string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Members = remove(d.Members, Member{ID: id})
}

func (d *DataStore) GetMembers() []Member {
	return d.Members
}

func (d *DataStore) GetMember(id string) *Member {
	i := find(d.Members, id)
	if i == -1 {
		return nil
	}
	return &d.Members[i]
}
func (d *DataStore) UpdateMember(m Member) *Member {
	d.mu.Lock()
	defer d.mu.Unlock()
	i := find(d.Members, m.ID)
	if i == -1 {
		return nil
	}
	d.Members[i] = m
	return &d.Members[i]
}

func (d *DataStore) Seed() {
	d.Members = []Member{
		{ID: "1", Name: "Alice"},
		{ID: "2", Name: "Bob"},
		{ID: "3", Name: "Sarah"},
		{ID: "4", Name: "Nash"},
		{ID: "5", Name: "Jhoe"},
	}
	d.Year = time.Now().Year()
	d.Exchanges = make(map[string]string)
}

// move this to utils
func remove(l []Member, item Member) []Member {
	for i, other := range l {
		if other.ID == item.ID {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

func find(l []Member, item string) int {
	for i, m := range l {
		if m.ID == item {
			return i
		}
	}
	return -1
}

func GetDS() *DataStore {
	return ds
}

func (d *DataStore) AssignRecipients() error {
	if len(d.Members) < 2 {
		return fmt.Errorf("at least two members are required")
	}

	givers := d.Members
	receivers := make([]Member, len(givers))
	copy(receivers, givers)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(receivers), func(i, j int) { receivers[i], receivers[j] = receivers[j], receivers[i] })

	for _, giver := range givers {
		assigned := false

		for i, receiver := range receivers {
			if receiver.ID == giver.ID {
				continue // Skip if the giver would give a gift to themselves.
			}

			if !hasGivenGiftRecently(giver, receiver) {
				d.Exchanges[giver.ID] = receiver.ID

				d.updateHistory(giver.ID, receiver.ID)

				// Remove the assigned receiver from the available receivers.
				receivers = append(receivers[:i], receivers[i+1:]...)
				assigned = true
				break
			}
		}

		if !assigned {
			return fmt.Errorf("unable to make valid exchange")
		}
	}

	return nil
}

func (d *DataStore) updateHistory(giverId, receiverId string) {
	for i, member := range d.Members {
		if member.ID == giverId {
			if len(d.Members[i].History) >= 3 {
				d.Members[i].History = d.Members[i].History[1:]
			}
			d.Members[i].History = append(d.Members[i].History, receiverId)
			break
		}
	}
}

func hasGivenGiftRecently(giver Member, receiver Member) bool {
	for _, recipient := range giver.History {
		if recipient == receiver.ID {
			return true
		}
	}
	return false
}
func (d *DataStore) GetExchanges() []Exchange {
	var xchanged []Exchange
	for k, v := range d.Exchanges {
		xchanged = append(xchanged, Exchange{Member: k, Recipient: v})
	}
	return xchanged
}
