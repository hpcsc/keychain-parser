package item

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFrom(t *testing.T) {
	t.Run("ignore inputs not starting with `keychain`", func(t *testing.T) {
		items := From([]string{
			"first",
			"second",
			"",
			"third",
		})

		require.Empty(t, items)
	})

	t.Run("return one item for each line starting with `keychain`", func(t *testing.T) {
		items := From([]string{
			"before",
			"keychain: \"key-chain-1\"",
			`class: "genp"`,
			"keychain: \"key-chain-2\"",
			`class: "genp"`,
			"keychain: \"key-chain-3\"",
			`class: "genp"`,
		})

		require.Len(t, items, 3)
	})

	t.Run("parse `svce` as service", func(t *testing.T) {
		items := From([]string{
			`keychain: "key-chain-1"`,
			`class: "genp"`,
			`	"svce"<blob>="service-1"`,
		})

		require.Len(t, items, 1)
		require.Equal(t, "service-1", items[0].Service)
	})

	t.Run("parse `acct` as account", func(t *testing.T) {
		items := From([]string{
			`keychain: "key-chain-1"`,
			`class: "genp"`,
			`	"acct"<blob>="account-1"`,
		})

		require.Len(t, items, 1)
		require.Equal(t, "account-1", items[0].Account)
	})

	t.Run("parse `gena` as attribute", func(t *testing.T) {
		items := From([]string{
			`keychain: "key-chain-1"`,
			`class: "genp"`,
			`	"gena"<blob>="attribute-1"`,
		})

		require.Len(t, items, 1)
		require.Equal(t, "attribute-1", items[0].Attribute)
	})

	t.Run("parse `icmt` as comment", func(t *testing.T) {
		items := From([]string{
			`keychain: "key-chain-1"`,
			`class: "genp"`,
			`	"icmt"<blob>="comment-1"`,
		})

		require.Len(t, items, 1)
		require.Equal(t, "comment-1", items[0].Comment)
	})

	t.Run("parse `0x00000007` as label", func(t *testing.T) {
		items := From([]string{
			`keychain: "key-chain-1"`,
			`class: "genp"`,
			`	0x00000007 <blob>="label-1"`,
		})

		require.Len(t, items, 1)
		require.Equal(t, "label-1", items[0].Label)
	})

	t.Run("parse class correctly", func(t *testing.T) {
		items := From([]string{
			`keychain: "key-chain-1"`,
			`class: "genp"`,
		})

		require.Len(t, items, 1)
		require.Equal(t, "genp", items[0].Class)
	})

	t.Run("only parse known classes", func(t *testing.T) {
		items := From([]string{
			`keychain: "key-chain-1"`,
			`class: "genp"`,
			`keychain: "key-chain-2"`,
			`class: "inet"`,
			`keychain: "key-chain-3"`,
			`class: "cert"`,
			`keychain: "key-chain-4"`,
			`class: "key"`,
			`keychain: "key-chain-5"`,
			`class: "unknown"`,
		})

		require.Len(t, items, 4)
		var classes []string
		for _, i := range items {
			classes = append(classes, i.Class)
		}
		require.ElementsMatch(t, []string{
			"genp",
			"inet",
			"cert",
			"key",
		}, classes)
	})
}
