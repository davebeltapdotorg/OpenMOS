package xml

import (
	"encoding/xml"
)

// UnmarshalXML accepts flat items and items with their fields inside mosItem.
// Nonempty outer fields win; a path container is selected as a whole.
// ponytail: abstracts are decoded as text; preserve nested markup if peer evidence requires it.
func (i *ItemInfo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw struct {
		ID                  string                `xml:"itemID"`
		Slug                string                `xml:"itemSlug"`
		Abstract            string                `xml:"mosAbstract"`
		Duration            string                `xml:"itemEdDur"`
		ObjDur              string                `xml:"objDur"`
		ObjTB               string                `xml:"objTB"`
		ObjectID            string                `xml:"objID"`
		MosID               string                `xml:"mosID"`
		ObjPath             string                `xml:"objPath"`
		ObjPaths            *ObjPaths             `xml:"objPaths"`
		Channel             string                `xml:"itemChannel"`
		MosExternalMetadata []MosExternalMetadata `xml:"mosExternalMetadata"`
		Nested              *struct {
			ID                  string                `xml:"itemID"`
			Slug                string                `xml:"itemSlug"`
			Abstract            string                `xml:"mosAbstract"`
			Duration            string                `xml:"itemEdDur"`
			ObjDur              string                `xml:"objDur"`
			ObjTB               string                `xml:"objTB"`
			ObjectID            string                `xml:"objID"`
			MosID               string                `xml:"mosID"`
			ObjPath             string                `xml:"objPath"`
			ObjPaths            *ObjPaths             `xml:"objPaths"`
			Channel             string                `xml:"itemChannel"`
			MosExternalMetadata []MosExternalMetadata `xml:"mosExternalMetadata"`
		} `xml:"mosItem"`
	}
	if err := d.DecodeElement(&raw, &start); err != nil {
		return err
	}

	i.ID, i.Slug, i.Abstract = raw.ID, raw.Slug, raw.Abstract
	i.Duration, i.ObjDur, i.ObjTB = raw.Duration, raw.ObjDur, raw.ObjTB
	i.ObjectID, i.MosID, i.ObjPath = raw.ObjectID, raw.MosID, raw.ObjPath
	i.ObjPaths, i.Channel, i.MosExternalMetadata = raw.ObjPaths, raw.Channel, raw.MosExternalMetadata
	if n := raw.Nested; n != nil {
		if i.ID == "" {
			i.ID = n.ID
		}
		if i.Slug == "" {
			i.Slug = n.Slug
		}
		if i.Abstract == "" {
			i.Abstract = n.Abstract
		}
		if i.Duration == "" {
			i.Duration = n.Duration
		}
		if i.ObjDur == "" {
			i.ObjDur = n.ObjDur
		}
		if i.ObjTB == "" {
			i.ObjTB = n.ObjTB
		}
		if i.ObjectID == "" {
			i.ObjectID = n.ObjectID
		}
		if i.MosID == "" {
			i.MosID = n.MosID
		}
		if i.ObjPath == "" {
			i.ObjPath = n.ObjPath
		}
		if i.ObjPaths.Empty() {
			i.ObjPaths = n.ObjPaths
		}
		if i.Channel == "" {
			i.Channel = n.Channel
		}
		if len(i.MosExternalMetadata) == 0 {
			i.MosExternalMetadata = n.MosExternalMetadata
		}
	}
	return nil
}
