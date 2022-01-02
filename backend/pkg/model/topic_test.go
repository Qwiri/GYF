package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopicArray_Add(t *testing.T) {
	topicA, topicB := NewTopic("hello"), NewTopic("world")
	topics := TopicArray{topicA, topicB}

	topics.Add("hi")

	assert.Equal(t, 3, len(topics))
	assert.Equal(t, topics[0], topicA)
	assert.Equal(t, topics[1], topicB)
	assert.EqualValues(t, topics[2], NewTopic("hi"))

	topics.Add(":)")

	assert.Equal(t, 4, len(topics))
	assert.Equal(t, topics[0], topicA)
	assert.Equal(t, topics[1], topicB)
	assert.EqualValues(t, topics[2], NewTopic("hi"))
	assert.EqualValues(t, topics[3], NewTopic(":)"))
}

func TestTopicArray_Delete(t *testing.T) {
	topicA, topicB := NewTopic("hello"), NewTopic("world")
	topics := TopicArray{topicA, topicB}

	topics.Delete("hello")

	assert.Equal(t, 1, len(topics))
	assert.Equal(t, topics[0], topicB)
}

func TestTopicArray_Exists(t *testing.T) {
	topics := TopicArray{
		NewTopic("hello"),
		NewTopic("world"),
		NewTopic("hi"),
		NewTopic(":)"),
	}
	assert.True(t, topics.Exists("hello"))
	assert.True(t, topics.Exists("HeLlO"))
	assert.True(t, topics.Exists("world"))
	assert.True(t, topics.Exists("WoRlD"))
	assert.True(t, topics.Exists("hi"))
	assert.True(t, topics.Exists("Hi"))
	assert.True(t, topics.Exists(":)"))
	assert.False(t, topics.Exists("sap?"))
}

func TestTopicArray_NextTopic(t *testing.T) {
	topics := TopicArray{
		NewTopic("hello"),
		NewTopic("world"),
	}

	next, err := topics.NextTopic()
	assert.Nil(t, err)
	assert.EqualValues(t, topics[0], next)
	next.Played = true

	next, err = topics.NextTopic()
	assert.Nil(t, err)
	assert.EqualValues(t, topics[1], next)
	next.Played = true

	next, err = topics.NextTopic()
	assert.Error(t, err)
	assert.Nil(t, next)
}

func TestTopicArray_PlayedTopicsCount(t *testing.T) {
	topics := TopicArray{
		NewTopic("hello"),
		NewTopic("world"),
		NewTopic("hi"),
	}

	topics[0].Played = true
	assert.Equal(t, 1, topics.PlayedTopicsCount())

	topics[1].Played = true
	assert.Equal(t, 2, topics.PlayedTopicsCount())
}
