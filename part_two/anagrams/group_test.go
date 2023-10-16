package anagrams

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "two_pairs",
			args: args{
				in: []string{"a", "ab", "cd", "ba", "dc"},
			},
			want: [][]string{{"a"}, {"ab", "ba"}, {"cd", "dc"}},
		},
		{
			name: "no_groups",
			args: args{
				in: []string{"ab", "cd", "ef", "gh", "ik"},
			},
			want: [][]string{{"ab"}, {"cd"}, {"ef"}, {"gh"}, {"ik"}},
		},
		{
			name: "one_group_three_elems",
			args: args{
				in: []string{"abc", "cba", "bca"},
			},
			want: [][]string{{"abc", "cba", "bca"}},
		},
		{
			name: "two_groups_four_elems",
			args: args{
				in: []string{"abc", "cba", "bca", "bac", "bcd", "dcb", "cbd", "cdb"},
			},
			want: [][]string{{"abc", "cba", "bca", "bac"}, {"bcd", "dcb", "cbd", "cdb"}},
		},
		{
			name: "task_requirements",
			args: args{
				in: []string{"affx", "a", "ab", "ba", "nnx", "xnn", "cde", "edc", "dce", "xffa"},
			},
			want: [][]string{{"a"}, {"ab", "ba"}, {"nnx", "xnn"}, {"cde", "edc", "dce"}, {"affx", "xffa"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupAnagrams(tt.args.in)
			for _, want := range tt.want {
				assert.Contains(t, got, want)
			}
		})
	}
}
