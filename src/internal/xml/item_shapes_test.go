package xml

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestItemInfoChoosesFlatScalarsAndWholePathContainer(t *testing.T) {
	const wire = `<roCreate><roID>RO</roID><roSlug>Order</roSlug><story><storyID>ST</storyID><storySlug>Story</storySlug>
<item><itemID>OUTER</itemID><mosAbstract>Outer abstract</mosAbstract><objTB>50</objTB>
<objPaths><objPath techDescription="essence">https://example.test/essence</objPath></objPaths>
<mosItem><itemID>INNER</itemID><itemSlug>Nested slug</itemSlug><objDur>150</objDur><objTB>59.94</objTB>
<objPaths><objProxyPath>https://example.test/proxy</objProxyPath></objPaths>
<mosExternalMetadata><mosSchema>urn:example:item</mosSchema><mosPayload><value>opaque</value></mosPayload></mosExternalMetadata>
</mosItem></item></story></roCreate>`
	var ro RunningOrderInfo
	if err := xml.Unmarshal([]byte(wire), &ro); err != nil {
		t.Fatal(err)
	}
	item := ro.Stories[0].Items[0]
	if item.ID != "OUTER" || item.Slug != "Nested slug" || item.Abstract != "Outer abstract" || item.ObjDur != "150" || item.ObjTB != "50" {
		t.Fatalf("scalar precedence or nested fallback failed: %+v", item)
	}
	if item.ObjPaths == nil || len(item.ObjPaths.Essence) != 1 || len(item.ObjPaths.Proxy) != 0 || item.ObjPaths.Essence[0].Value != "https://example.test/essence" {
		t.Fatalf("path containers were merged or dropped: %+v", item.ObjPaths)
	}
	if len(item.MosExternalMetadata) != 1 || item.MosExternalMetadata[0].MosPayload.Raw != `<value>opaque</value>` {
		t.Fatalf("nested opaque metadata was lost: %+v", item.MosExternalMetadata)
	}
}

func TestItemInfoNestedOnlyKeepsAbstractSeparateFromSlug(t *testing.T) {
	var item ItemInfo
	if err := xml.Unmarshal([]byte(`<item><mosItem><itemID>I</itemID><objID>O</objID><mosID>openmos.example.test</mosID><mosAbstract> Full text </mosAbstract><objPaths><objProxyPath>https://example.test/proxy</objProxyPath></objPaths></mosItem></item>`), &item); err != nil {
		t.Fatal(err)
	}
	if item.ID != "I" || item.ObjectID != "O" || item.MosID != "openmos.example.test" || item.Slug != "" || item.Abstract != " Full text " {
		t.Fatalf("nested item fields were lost: %+v", item)
	}
	if item.ObjPaths == nil || len(item.ObjPaths.Proxy) != 1 {
		t.Fatalf("nested paths were lost: %+v", item.ObjPaths)
	}
}

func TestItemInfoDoesNotEmitInputOnlyFields(t *testing.T) {
	encoded, err := xml.Marshal(ItemInfo{ID: "I", ObjectID: "O", Abstract: "Full text", ObjDur: "150", ObjTB: "50", ObjPaths: &ObjPaths{Essence: []ObjPath{{Value: "https://example.test/essence"}}}})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(encoded), "mosAbstract") || strings.Contains(string(encoded), "objDur") || strings.Contains(string(encoded), "objTB") || strings.Contains(string(encoded), "objPaths") {
		t.Fatalf("input-only item fields emitted: %s", encoded)
	}
}
