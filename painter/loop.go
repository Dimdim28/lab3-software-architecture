package painter

import (
	"image"
	"sync"

	"golang.org/x/exp/shiny/screen"
)

// Receiver отримує текстуру, яка була підготовлена в результаті виконання команд у циелі подій.
type Receiver interface {
	Update(t screen.Texture)
}

// Loop реалізує цикл подій для формування текстури отриманої через виконання операцій отриманих з внутрішньої черги.
type Loop struct {
	Receiver Receiver

	next screen.Texture // текстура, яка зараз формується
	prev screen.Texture // текстура, яка була відправленя останнього разу у Receiver

	stopReq bool
	stopped chan struct{}

	MsgQueue messageQueue
}

var size = image.Pt(800, 800)

// Start запускає цикл подій. Цей метод потрібно запустити до того, як викликати на ньому будь-які інші методи.
func (l *Loop) Start(s screen.Screen) {
	l.next, _ = s.NewTexture(size)
	l.prev, _ = s.NewTexture(size)

	l.MsgQueue = messageQueue{}
	go l.eventProcess()
}

func (l *Loop) eventProcess() {
	for {
		if op := l.MsgQueue.Pull(); op != nil {
			if update := op.Do(l.next); update {
				l.Receiver.Update(l.next)
				l.next, l.prev = l.prev, l.next
			}
		}
	}
}

// Post додає нову операцію у внутрішню чергу.
func (l *Loop) Post(op Operation) {
	// TODO: реалізувати додавання операції в чергу. Поточна імплементація
	if op != nil {
		l.MsgQueue.Push(op)
	}
}

// StopAndWait сигналізує
func (l *Loop) StopAndWait() {
	l.Post(OperationFunc(func(screen.Texture) {
		l.stopReq = true
	}))
	<-l.stopped
}

// TODO: реалізувати власну чергу повідомлень.
type messageQueue struct {
	Queue   []Operation
	mu      sync.Mutex
	blocked chan struct{}
}

func (MsgQueue *messageQueue) Push(op Operation) {
	MsgQueue.mu.Lock()
	defer MsgQueue.mu.Unlock()
	MsgQueue.Queue = append(MsgQueue.Queue, op)
	if MsgQueue.blocked != nil {
		close(MsgQueue.blocked)
		MsgQueue.blocked = nil
	}
}

func (MsgQueue *messageQueue) Pull() Operation {
	MsgQueue.mu.Lock()
	defer MsgQueue.mu.Unlock()
	for len(MsgQueue.Queue) == 0 {
		MsgQueue.blocked = make(chan struct{})
		MsgQueue.mu.Unlock()
		<-MsgQueue.blocked
		MsgQueue.mu.Lock()
	}

	op := MsgQueue.Queue[0]
	MsgQueue.Queue[0] = nil
	MsgQueue.Queue = MsgQueue.Queue[1:]
	return op
}
