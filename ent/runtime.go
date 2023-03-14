// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/kzmijak/zswod_api_go/ent/article"
	"github.com/kzmijak/zswod_api_go/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	articleFields := schema.Article{}.Fields()
	_ = articleFields
	// articleDescTitle is the schema descriptor for title field.
	articleDescTitle := articleFields[1].Descriptor()
	// article.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	article.TitleValidator = func() func(string) error {
		validators := articleDescTitle.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(title string) error {
			for _, fn := range fns {
				if err := fn(title); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// articleDescShort is the schema descriptor for short field.
	articleDescShort := articleFields[3].Descriptor()
	// article.ShortValidator is a validator for the "short" field. It is called by the builders before save.
	article.ShortValidator = func() func(string) error {
		validators := articleDescShort.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(short string) error {
			for _, fn := range fns {
				if err := fn(short); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
