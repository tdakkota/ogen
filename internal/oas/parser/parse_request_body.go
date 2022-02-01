package parser

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/internal/oas"
)

func (p *parser) parseRequestBody(body *ogen.RequestBody) (*oas.RequestBody, error) {
	if ref := body.Ref; ref != "" {
		reqBody, err := p.resolveRequestBody(ref)
		if err != nil {
			return nil, errors.Wrapf(err, "resolve %q reference", ref)
		}

		return reqBody, nil
	}

	result := &oas.RequestBody{
		Contents: make(map[string]*oas.Schema, len(body.Content)),
		Required: body.Required,
	}

	for contentType, media := range body.Content {
		schema, err := p.parseMedia(media)
		if err != nil {
			return nil, errors.Wrapf(err, "content: %s", contentType)
		}
		result.Contents[contentType] = schema
	}

	return result, nil
}

func (p *parser) parseMedia(media ogen.Media) (*oas.Schema, error) {
	schema, err := p.parseSchema(media.Schema)
	if err != nil {
		return nil, errors.Wrap(err, "parse schema")
	}

	if m := media.Encoding; len(m) > 0 {
		for propName, encoding := range m {
			var prop *oas.Property
			for _, prop = range schema.Properties {
				if prop.Name == propName {
					break
				}
			}
			if prop == nil {
				return nil, errors.Errorf("unknown property %q", propName)
			}

			if err := p.parsePropertyEncoding(encoding, prop); err != nil {
				return nil, errors.Wrapf(err, "encoding: %q", propName)
			}
		}
	}
	return schema, nil
}

func (p *parser) parsePropertyEncoding(encoding ogen.Encoding, prop *oas.Property) error {
	locatedIn := oas.LocationQuery
	style, err := paramStyle(locatedIn, encoding.Style)
	if err != nil {
		return errors.Wrap(err, "style")
	}
	prop.Style = style

	prop.Explode = paramExplode(locatedIn, encoding.Explode)
	return nil
}
