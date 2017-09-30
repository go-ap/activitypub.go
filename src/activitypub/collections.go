package activitypub

var validCollectionTypes = [...]string{}

type Page ObjectOrLink

type Collection struct {
	*Object
	// A non-negative integer specifying the total number of objects contained by the logical view of the collection.
	// This number might not reflect the actual number of items serialized within the Collection object instance.
	TotalItems uint `jsonld:"totalItems"`
	// Identifies the items contained in a collection. The items might be ordered or unordered.
	Items ItemCollection `jsonld:"items"`
}

type OrderedCollection struct {
	*Object
	// A non-negative integer specifying the total number of objects contained by the logical view of the collection.
	// This number might not reflect the actual number of items serialized within the Collection object instance.
	TotalItems uint `jsonld:"totalItems"`
	// Identifies the items contained in a collection. The items might be ordered or unordered.
	OrderedItems ItemCollection `jsonld:"orderedItems"`
}

type CollectionPage struct {
	PartOf *Collection
	// In a paged Collection, indicates the page that contains the most recently updated member items.
	Current Page `jsonld:"current"`
	// In a paged Collection, indicates the furthest preceeding page of items in the collection.
	First Page `jsonld:"first"`
	// In a paged Collection, indicates the furthest proceeding page of the collection.
	Last Page `jsonld:"last"`
	// In a paged Collection, indicates the next page of items.
	Next Page `jsonld:"next"`
	// In a paged Collection, identifies the previous page of items.
	Prev Page `jsonld:"prev"`
}

type OrderedCollectionPage struct {
	PartOf *OrderedCollection
	// In a paged Collection, indicates the page that contains the most recently updated member items.
	Current Page `jsonld:"current"`
	// In a paged Collection, indicates the furthest preceeding page of items in the collection.
	First Page `jsonld:"first"`
	// In a paged Collection, indicates the furthest proceeding page of the collection.
	Last Page `jsonld:"last"`
	// In a paged Collection, indicates the next page of items.
	Next Page `jsonld:"next"`
	// In a paged Collection, identifies the previous page of items.
	Prev Page `jsonld:"prev"`
	// A non-negative integer value identifying the relative position within the logical view of a strictly ordered collection.
	StartIndex uint `jsonld:"startIndex"`
}

func ValidCollectionType(_type string) bool {
	for _, v := range validCollectionTypes {
		if v == _type {
			return true
		}
	}
	return false
}

func CollectionNew(id ObjectId) *Collection {
	o := ObjectNew(id, CollectionType)

	return &Collection{Object: o}
}

func OrderedCollectionNew(id ObjectId) *OrderedCollection {
	o := ObjectNew(id, OrderedCollectionType)

	return &OrderedCollection{Object: o}
}

func CollectionPageNew(parent *Collection) *CollectionPage {
	return &CollectionPage{PartOf: parent}
}
func OrderedCollectionPageNew(parent *OrderedCollection) *OrderedCollectionPage {
	return &OrderedCollectionPage{PartOf: parent}
}