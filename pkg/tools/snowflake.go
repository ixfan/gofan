package tools

import "github.com/bwmarrin/snowflake"

//生成新ID
var snowFlakeNode *snowflake.Node

func NewSnowflakeId() (int64, error) {
	if snowFlakeNode == nil {
		node, err := snowflake.NewNode(1)
		if err != nil {
			return 0, err
		}
		snowFlakeNode = node
	}
	// Generate a snowflake ID.
	id := snowFlakeNode.Generate()
	return id.Int64(), nil
}
