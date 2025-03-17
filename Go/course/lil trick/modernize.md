# Analyzer modernize ¶

简化代码：通过使用Go的现代构造来简化和澄清现有代码

此分析器会报告通过使用Go的更现代功能来简化和澄清现有代码的机会，例如：

- 使用Go 1.21中添加的内置min或max函数代替if/else条件赋值；
- 使用Go 1.21中添加的slices.Sort(s)代替sort.Slice(x, func(i, j int) bool) { return s[i] < s[j] }；
- 使用Go 1.18中添加的'any'类型代替interface{}；
- 使用Go 1.21中添加的slices.Clone(s)或slices.Concat(s)代替append([]T(nil), s...)；
- 使用Go 1.21中maps包添加的Collect、Copy、Clone或Insert函数代替对map的m[k]=v更新循环；
- 使用Go 1.19中添加的fmt.Appendf(nil, ...)代替[]byte(fmt.Sprintf(...))；
- 使用Go 1.24中添加的t.Context代替测试中的context.WithCancel；
- 在结构体中使用omitzero代替omitempty，Go 1.24中添加；
- 使用Go 1.21中添加的slices.Delete(s, i, i+1)代替append(s[:i], s[i+1]...)；
- 使用Go 1.22中添加的for i := range n {}代替3-clause的for i := 0; i < n; i++ {}循环；
- 在"for range strings.Split(...)"中使用Go 1.24中更高效的SplitSeq代替Split；
要批量应用所有现代化修复，可以使用以下命令：

`$ go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest -test ./...`

---

modernize: simplify code by using modern constructs

This analyzer reports opportunities for simplifying and clarifying existing code by using more modern features of Go, such as:

replacing an if/else conditional assignment by a call to the built-in min or max functions added in go1.21;
replacing sort.Slice(x, func(i, j int) bool) { return s[i] < s[j] } by a call to slices.Sort(s), added in go1.21;
replacing interface{} by the 'any' type added in go1.18;
replacing append([]T(nil), s...) by slices.Clone(s) or slices.Concat(s), added in go1.21;
replacing a loop around an m[k]=v map update by a call to one of the Collect, Copy, Clone, or Insert functions from the maps package, added in go1.21;
replacing []byte(fmt.Sprintf...) by fmt.Appendf(nil, ...), added in go1.19;
replacing uses of context.WithCancel in tests with t.Context, added in go1.24;
replacing omitempty by omitzero on structs, added in go1.24;
replacing append(s[:i], s[i+1]...) by slices.Delete(s, i, i+1), added in go1.21
replacing a 3-clause for i := 0; i < n; i++ {} loop by for i := range n {}, added in go1.22;
replacing Split in "for range strings.Split(...)" by go1.24's more efficient SplitSeq;
To apply all modernization fixes en masse, you can use the following command:

`$ go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest -test ./...`
