package main

import "fmt"

// Command Интерфейс команды
type Command interface {
	Execute()
}

// LightOnCommand Конкретная команда
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

// RemoteControl объект, выполняющий команды
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

// Light представляет устройство (в данном случае, свет)
type Light struct {
	on bool
}

func (l *Light) On() {
	l.on = true
	fmt.Println("Light is on")
}

func (l *Light) Off() {
	l.on = false
	fmt.Println("Light is off")
}

func main() {
	// Создаем объекты команды и устройства
	light := &Light{}
	lightOnCommand := &LightOnCommand{light: light}

	// Создаем пульт управления и устанавливаем команду
	remoteControl := &RemoteControl{}
	remoteControl.SetCommand(lightOnCommand)

	// Нажимаем кнопку на пульте
	remoteControl.PressButton()
}
