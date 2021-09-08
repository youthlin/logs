package logs

import "github.com/youthlin/logs/pkg/trie"

func LevelConfig(lvl Level) *LoggerLevel {
	Assert(lvl > LevelUnset, "root level must set")
	return &LoggerLevel{Root: lvl}
}

func (c *LoggerLevel) Trie() *trie.Tire {
	if c.trie == nil {
		root := trie.NewTire(c.Root)
		for name, lvl := range c.Loggers {
			if lvl != LevelUnset {
				root.Insert(name, lvl)
			}
		}
		c.trie = root
	}
	return c.trie
}
