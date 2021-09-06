package logs

import "github.com/youthlin/logs/pkg/trie"

func LevelConfig(lvl Level) *Config {
	return &Config{Root: lvl}
}

func (c *Config) Trie() *trie.Tire {
	if c.trie == nil {
		root := trie.NewTire(c.Root)
		for name, lvl := range c.Loggers {
			root.Insert(name, lvl)
		}
		c.trie = root
	}
	return c.trie
}
