//nolint:gochecknoglobals // it's fine
package otherworld

import "github.com/limike954/trajectory-data-00009/pkg/cardinal"

// Matchmaking is another shard. Just for example send this to itself.
var Matchmaking = cardinal.OtherWorld{
	Region:       "us-west-2",
	Organization: "organization",
	Project:      "project",
	ShardID:      "game", // The shard ID of the other shard.
}
