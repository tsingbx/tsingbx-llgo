; ModuleID = 'main'
source_filename = "main"

%"github.com/goplus/llgo/internal/runtime.String" = type { ptr, i64 }
%"github.com/goplus/llgo/internal/runtime.Slice" = type { ptr, i64, i64 }

@"main.init$guard" = global ptr null
@0 = private unnamed_addr constant [1 x i8] zeroinitializer, align 1
@1 = private unnamed_addr constant [6 x i8] c"Hello\00", align 1
@2 = private unnamed_addr constant [2 x i8] c" \00", align 1
@3 = private unnamed_addr constant [6 x i8] c"World\00", align 1
@__stderrp = external global ptr
@4 = private unnamed_addr constant [10 x i8] c"Hello %s\0A\00", align 1

define %"github.com/goplus/llgo/internal/runtime.String" @main.concat(%"github.com/goplus/llgo/internal/runtime.Slice" %0) {
_llgo_0:
  %1 = call i64 @"github.com/goplus/llgo/internal/runtime.SliceLen"(%"github.com/goplus/llgo/internal/runtime.Slice" %0)
  br label %_llgo_1

_llgo_1:                                          ; preds = %_llgo_2, %_llgo_0
  %2 = phi %"github.com/goplus/llgo/internal/runtime.String" [ %10, %_llgo_0 ], [ %9, %_llgo_2 ]
  %3 = phi i64 [ -1, %_llgo_0 ], [ %4, %_llgo_2 ]
  %4 = add i64 %3, 1
  %5 = icmp slt i64 %4, %1
  br i1 %5, label %_llgo_2, label %_llgo_3

_llgo_2:                                          ; preds = %_llgo_1
  %6 = call ptr @"github.com/goplus/llgo/internal/runtime.SliceData"(%"github.com/goplus/llgo/internal/runtime.Slice" %0)
  %7 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %6, i64 %4
  %8 = load %"github.com/goplus/llgo/internal/runtime.String", ptr %7, align 8
  %9 = call %"github.com/goplus/llgo/internal/runtime.String" @"github.com/goplus/llgo/internal/runtime.StringCat"(%"github.com/goplus/llgo/internal/runtime.String" %2, %"github.com/goplus/llgo/internal/runtime.String" %8)
  br label %_llgo_1

_llgo_3:                                          ; preds = %_llgo_1
  ret %"github.com/goplus/llgo/internal/runtime.String" %2
  %10 = call %"github.com/goplus/llgo/internal/runtime.String" @"github.com/goplus/llgo/internal/runtime.NewString"(ptr @0, i64 0)
}

define void @main.init() {
_llgo_0:
  %0 = load i1, ptr @"main.init$guard", align 1
  br i1 %0, label %_llgo_2, label %_llgo_1

_llgo_1:                                          ; preds = %_llgo_0
  store i1 true, ptr @"main.init$guard", align 1
  br label %_llgo_2

_llgo_2:                                          ; preds = %_llgo_1, %_llgo_0
  ret void
}

define void @main() {
_llgo_0:
  call void @"github.com/goplus/llgo/internal/runtime.init"()
  call void @main.init()
  %0 = call ptr @"github.com/goplus/llgo/internal/runtime.Alloc"(i64 48)
  %1 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %0, i64 0
  %2 = call %"github.com/goplus/llgo/internal/runtime.String" @"github.com/goplus/llgo/internal/runtime.NewString"(ptr @1, i64 5)
  store %"github.com/goplus/llgo/internal/runtime.String" %2, ptr %1, align 8
  %3 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %0, i64 1
  %4 = call %"github.com/goplus/llgo/internal/runtime.String" @"github.com/goplus/llgo/internal/runtime.NewString"(ptr @2, i64 1)
  store %"github.com/goplus/llgo/internal/runtime.String" %4, ptr %3, align 8
  %5 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %0, i64 2
  %6 = call %"github.com/goplus/llgo/internal/runtime.String" @"github.com/goplus/llgo/internal/runtime.NewString"(ptr @3, i64 5)
  store %"github.com/goplus/llgo/internal/runtime.String" %6, ptr %5, align 8
  %7 = call %"github.com/goplus/llgo/internal/runtime.Slice" @"github.com/goplus/llgo/internal/runtime.NewSlice"(ptr %0, i64 3, i64 3)
  %8 = call %"github.com/goplus/llgo/internal/runtime.String" @main.concat(%"github.com/goplus/llgo/internal/runtime.Slice" %7)
  %9 = load ptr, ptr @__stderrp, align 8
  %10 = call i64 @"github.com/goplus/llgo/internal/runtime.StringLen"(%"github.com/goplus/llgo/internal/runtime.String" %8)
  %11 = add i64 %10, 1
  %12 = alloca i8, i64 %11, align 1
  %13 = call ptr @"github.com/goplus/llgo/internal/runtime.CStrCopy"(ptr %12, %"github.com/goplus/llgo/internal/runtime.String" %8)
  %14 = call i32 (ptr, ptr, ...) @fprintf(ptr %9, ptr @4, ptr %13)
  ret void
}

declare i64 @"github.com/goplus/llgo/internal/runtime.SliceLen"(%"github.com/goplus/llgo/internal/runtime.Slice")

declare ptr @"github.com/goplus/llgo/internal/runtime.SliceData"(%"github.com/goplus/llgo/internal/runtime.Slice")

declare %"github.com/goplus/llgo/internal/runtime.String" @"github.com/goplus/llgo/internal/runtime.StringCat"(%"github.com/goplus/llgo/internal/runtime.Slice")

declare %"github.com/goplus/llgo/internal/runtime.String" @"github.com/goplus/llgo/internal/runtime.NewString"(ptr, i64)

declare void @"github.com/goplus/llgo/internal/runtime.init"()

declare ptr @"github.com/goplus/llgo/internal/runtime.Alloc"(i64)

declare %"github.com/goplus/llgo/internal/runtime.Slice" @"github.com/goplus/llgo/internal/runtime.NewSlice"(ptr, i64, i64)

declare i64 @"github.com/goplus/llgo/internal/runtime.StringLen"(%"github.com/goplus/llgo/internal/runtime.Slice")

declare ptr @"github.com/goplus/llgo/internal/runtime.CStrCopy"(ptr, %"github.com/goplus/llgo/internal/runtime.String")

declare i32 @fprintf(ptr, ptr, ...)