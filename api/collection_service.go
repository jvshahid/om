package api

import (
	"fmt"
	"regexp"
	"sort"
)

type responsePropertyCollection []responsePropertyCollectionItem

type updatedPropertyCollectionItem struct {
	Data map[string]interface{}
}

type updatedPropertyCollection []updatedPropertyCollectionItem

func (i updatedPropertyCollectionItem) getFieldValue(fieldName string) string {
	if value, ok := i.Data[fieldName].(string); ok {
		return value
	}
	return ""
}

func (i updatedPropertyCollectionItem) setFieldValue(fieldName string, value string) {
	i.Data[fieldName] = value
}

func parseUpdatedPropertyCollection(updatedProperty interface{}) (updatedPropertyCollection, error) {
	var collection updatedPropertyCollection

	updatedPropertyAsMap, ok := updatedProperty.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("parseUpdatedPropertyCollection: failed to convert %v to map[string]interface{}", updatedProperty)
	}

	rawItems := updatedPropertyAsMap["value"]

	rawItemSlice, ok := rawItems.([]interface{})
	if !ok {
		return nil, fmt.Errorf("parseUpdatedPropertyCollection: failed to convert %v to []interface{}", rawItems)
	}

	for _, item := range rawItemSlice {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("parseUpdatedPropertyCollection: failed to convert %v to map[string]interface{}", item)
		}

		collection = append(collection, updatedPropertyCollectionItem{Data: itemMap})
	}

	return collection, nil
}

type responsePropertyCollectionItem struct {
	Data map[interface{}]interface{}
}

func parseResponsePropertyCollection(rawItems interface{}) (responsePropertyCollection, error) {
	var collection responsePropertyCollection
	if rawItemSlice, ok := rawItems.([]interface{}); ok {
		for _, item := range rawItemSlice {
			if itemMap, ok := item.(map[interface{}]interface{}); ok {
				collection = append(collection, responsePropertyCollectionItem{Data: itemMap})
			} else {
				return nil, fmt.Errorf("parseResponsePropertyCollection: failed to convert %v to map[interface{}]interface{}", item)
			}
		}
	} else {
		return nil, fmt.Errorf("parseResponsePropertyCollection: failed to convert %v to []interface{}", rawItems)
	}
	return collection, nil
}

func (i responsePropertyCollectionItem) getFieldValue(fieldName string) string {
	if valueObj, ok := i.Data[fieldName].(map[interface{}]interface{}); ok {
		if value, ok := valueObj["value"].(string); ok {
			return value
		}
	}
	return ""
}

func (i responsePropertyCollectionItem) findLogicalKeyField() string {
	sortedFields := sortFields(i)

	regexes := []string{"^name$", "^key$", "(?i)name$"}
	for _, regex := range regexes {
		compiledRegex := regexp.MustCompile(regex)

		for _, str := range sortedFields {
			if compiledRegex.MatchString(str) {
				return str
			}
		}
	}

	return ""
}

func sortFields(item responsePropertyCollectionItem) []string {
	sortedFields := make([]string, 0, len(item.Data))
	for k := range item.Data {
		sortedFields = append(sortedFields, fmt.Sprintf("%v", k))
	}
	sort.Strings(sortedFields)

	return sortedFields
}

func associateExistingCollectionGUIDs(updatedProperty interface{}, existingProperty ResponseProperty) error {
	existingCollection, err := parseResponsePropertyCollection(existingProperty.Value)
	if err != nil {
		return err
	}

	updatedCollection, err := parseUpdatedPropertyCollection(updatedProperty)
	if err != nil {
		return err
	}

	logicalKeyFieldName := existingCollection[0].findLogicalKeyField()

	assignExistingGUIDtoPropertyItem(existingCollection, updatedCollection, logicalKeyFieldName)

	return nil
}

func assignExistingGUIDtoPropertyItem(existingCollection responsePropertyCollection, updatedCollection updatedPropertyCollection, logicalKeyFieldName string) {
	for _, updatedCollectionItem := range updatedCollection {
		updatedCollectionItemLogicalKeyValue := updatedCollectionItem.getFieldValue(logicalKeyFieldName)
		for _, existingCollectionItem := range existingCollection {
			if updatedCollectionItemLogicalKeyValue == existingCollectionItem.getFieldValue(logicalKeyFieldName) {
				updatedCollectionItem.setFieldValue("guid", existingCollectionItem.getFieldValue("guid"))
				break //move onto the next updatedCollectionItem
			}
		}
	}
}
