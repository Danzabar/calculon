package main

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "math/rand"
)

type Library struct {
    Greetings  []string `yaml:"greetings"`
    Insults    []string `yaml:"insults"`
    Intros     []string `yaml:"intros"`
    Statements []string `yaml:"statements"`
    Confused   []string `yaml:"confused"`
}

func (l *Library) RandomGreeting() string {
    return l.Greetings[rand.Intn(len(l.Greetings))]
}

func (l *Library) RandomInsult() string {
    return l.Insults[rand.Intn(len(l.Insults))]
}

func (l *Library) RandomStatement() string {
    return l.Statements[rand.Intn(len(l.Statements))]
}

func (l *Library) RandomIntro() string {
    return l.Intros[rand.Intn(len(l.Intros))]
}

func (l *Library) RandomConfused() string {
    return l.Confused[rand.Intn(len(l.Confused))]
}

// Reads the YAML file and loads into a library
func SetupLibrary() *Library {
    l := &Library{}

    d, err := ioutil.ReadFile("./config/library.yml")

    if err != nil {
        log.Fatal(err)
    }

    if e := yaml.Unmarshal(d, l); e != nil {
        log.Fatal(err)
    }

    return l
}
