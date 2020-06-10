/*
Notify system
*/

package notify

import "context"

var (
	subsribers = Notify{
		subsribers: map[int][]Subscriber{},
	}
)

func Subscribe(event int, subscriber Subscriber) { // nolint unused
	subsribers.Lock()
	subsribers.subsribers[event] = append(subsribers.subsribers[event], subscriber)
	subsribers.Unlock()
}

func UnSubscribe(event int, subscriber Subscriber) { // nolint unused
	subsribers.Lock()
	defer subsribers.Unlock()

	for _, v := range subsribers.subsribers[event] {
		if subscriber == v {
			delete(subsribers.subsribers, event)
			break
		}
	}
}

func Publish(ctx context.Context, event int, payload interface{}, cb chan<- interface{}, responseFilter string) { // nolint unused
	responses := map[string]*Response{}
	subsribers.Lock()
	defer subsribers.Unlock()

	if len(subsribers.subsribers[event]) == 0 {
		cb <- nil
	}

	for key := range subsribers.subsribers[event] {
		response := subsribers.subsribers[event][key].Notify(ctx, event, payload)
		if response != nil && response.Error != nil {
			if response.Error != nil {
				// TODO: Need to add undo operations
				cb <- *response
				return
			}

			responses[response.Name] = response
		}
	}

	// TODO: Send only first success response for simple implementation
	if responses[responseFilter] != nil {
		cb <- *responses[responseFilter]
	}
}

func Clean() { // nolint unused
	subsribers = Notify{
		subsribers: map[int][]Subscriber{},
	}
}
