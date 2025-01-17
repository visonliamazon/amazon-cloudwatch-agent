// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package rules

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type ReplaceActions struct {
	Actions []ActionItem
}

func NewReplacer(rules []Rule) *ReplaceActions {
	return &ReplaceActions{
		generateActionDetails(rules, AllowListActionReplace),
	}
}

func (r *ReplaceActions) Process(attributes, _ pcommon.Map, isTrace bool) error {
	// do nothing when there is no replace rule defined
	if r.Actions == nil || len(r.Actions) == 0 {
		return nil
	}
	// If there are more than one rule are matched, the last one will be executed(Later one has higher priority)
	actions := r.Actions
	finalRules := make(map[string]string)
	for i := len(actions) - 1; i >= 0; i = i - 1 {
		element := actions[i]
		isMatched := matchesSelectors(attributes, element.SelectorMatchers, isTrace)
		if !isMatched {
			continue
		}
		for _, replacement := range element.Replacements {
			targetDimensionKey := getExactKey(replacement.TargetDimension, isTrace)
			// don't allow customer add new dimension key
			_, isExist := attributes.Get(targetDimensionKey)
			if !isExist {
				continue
			}
			// every replacement in one specific dimension only will be performed once
			_, ok := finalRules[targetDimensionKey]
			if ok {
				continue
			}
			finalRules[targetDimensionKey] = replacement.Value
		}
	}

	for key, value := range finalRules {
		attributes.PutStr(key, value)
	}
	return nil
}
