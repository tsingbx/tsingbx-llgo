; ModuleID = 'main'
source_filename = "main"

%main.bar = type { ptr, float }
%"github.com/goplus/llgo/internal/runtime.eface" = type { ptr, ptr }
%"github.com/goplus/llgo/internal/runtime.String" = type { ptr, i64 }
%"github.com/goplus/llgo/internal/abi.StructField" = type { %"github.com/goplus/llgo/internal/runtime.String", ptr, i64, %"github.com/goplus/llgo/internal/runtime.String", i1 }
%"github.com/goplus/llgo/internal/runtime.Slice" = type { ptr, i64, i64 }

@"main.init$guard" = global ptr null
@0 = private unnamed_addr constant [4 x i8] c"Foo\00", align 1
@main.bar = global ptr null
@1 = private unnamed_addr constant [9 x i8] c"Foo done\00", align 1
@__llgo_argc = global ptr null
@__llgo_argv = global ptr null
@2 = private unnamed_addr constant [5 x i8] c"main\00", align 1
@3 = private unnamed_addr constant [7 x i8] c"notOk:\00", align 1
@4 = private unnamed_addr constant [3 x i8] c"pb\00", align 1
@"*_llgo_byte" = linkonce global ptr null
@5 = private unnamed_addr constant [1 x i8] zeroinitializer, align 1
@6 = private unnamed_addr constant [2 x i8] c"f\00", align 1
@7 = private unnamed_addr constant [1 x i8] zeroinitializer, align 1
@8 = private unnamed_addr constant [5 x i8] c"main\00", align 1
@9 = private unnamed_addr constant [9 x i8] c"main.bar\00", align 1

define { %main.bar, i1 } @main.Foo(%"github.com/goplus/llgo/internal/runtime.eface" %0) {
_llgo_0:
  %1 = alloca %"github.com/goplus/llgo/internal/runtime.String", align 8
  %2 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %1, i32 0, i32 0
  store ptr @0, ptr %2, align 8
  %3 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %1, i32 0, i32 1
  store i64 3, ptr %3, align 4
  %4 = load %"github.com/goplus/llgo/internal/runtime.String", ptr %1, align 8
  call void @"github.com/goplus/llgo/internal/runtime.PrintString"(%"github.com/goplus/llgo/internal/runtime.String" %4)
  call void @"github.com/goplus/llgo/internal/runtime.PrintByte"(i8 10)
  %5 = extractvalue %"github.com/goplus/llgo/internal/runtime.eface" %0, 0
  %6 = load ptr, ptr @main.bar, align 8
  %7 = icmp eq ptr %5, %6
  br i1 %7, label %_llgo_1, label %_llgo_2

_llgo_1:                                          ; preds = %_llgo_0
  %8 = extractvalue %"github.com/goplus/llgo/internal/runtime.eface" %0, 1
  %9 = load %main.bar, ptr %8, align 8
  %10 = alloca { %main.bar, i1 }, align 8
  %11 = getelementptr inbounds { %main.bar, i1 }, ptr %10, i32 0, i32 0
  store %main.bar %9, ptr %11, align 8
  %12 = getelementptr inbounds { %main.bar, i1 }, ptr %10, i32 0, i32 1
  store i1 true, ptr %12, align 1
  %13 = load { %main.bar, i1 }, ptr %10, align 8
  br label %_llgo_3

_llgo_2:                                          ; preds = %_llgo_0
  %14 = alloca { %main.bar, i1 }, align 8
  %15 = getelementptr inbounds { %main.bar, i1 }, ptr %14, i32 0, i32 0
  store { ptr, double } zeroinitializer, ptr %15, align 8
  %16 = getelementptr inbounds { %main.bar, i1 }, ptr %14, i32 0, i32 1
  store i1 false, ptr %16, align 1
  %17 = load { %main.bar, i1 }, ptr %14, align 8
  br label %_llgo_3

_llgo_3:                                          ; preds = %_llgo_2, %_llgo_1
  %18 = phi { %main.bar, i1 } [ %13, %_llgo_1 ], [ %17, %_llgo_2 ]
  %19 = extractvalue { %main.bar, i1 } %18, 0
  %20 = extractvalue { %main.bar, i1 } %18, 1
  %21 = alloca %"github.com/goplus/llgo/internal/runtime.String", align 8
  %22 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %21, i32 0, i32 0
  store ptr @1, ptr %22, align 8
  %23 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %21, i32 0, i32 1
  store i64 8, ptr %23, align 4
  %24 = load %"github.com/goplus/llgo/internal/runtime.String", ptr %21, align 8
  call void @"github.com/goplus/llgo/internal/runtime.PrintString"(%"github.com/goplus/llgo/internal/runtime.String" %24)
  call void @"github.com/goplus/llgo/internal/runtime.PrintByte"(i8 10)
  %mrv = insertvalue { %main.bar, i1 } poison, %main.bar %19, 0
  %mrv1 = insertvalue { %main.bar, i1 } %mrv, i1 %20, 1
  ret { %main.bar, i1 } %mrv1
}

define void @main.init() {
_llgo_0:
  %0 = load i1, ptr @"main.init$guard", align 1
  br i1 %0, label %_llgo_2, label %_llgo_1

_llgo_1:                                          ; preds = %_llgo_0
  store i1 true, ptr @"main.init$guard", align 1
  call void @"main.init$abi"()
  br label %_llgo_2

_llgo_2:                                          ; preds = %_llgo_1, %_llgo_0
  ret void
}

define i32 @main(i32 %0, ptr %1) {
_llgo_0:
  store i32 %0, ptr @__llgo_argc, align 4
  store ptr %1, ptr @__llgo_argv, align 8
  call void @"github.com/goplus/llgo/internal/runtime.init"()
  call void @main.init()
  %2 = alloca %"github.com/goplus/llgo/internal/runtime.String", align 8
  %3 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %2, i32 0, i32 0
  store ptr @2, ptr %3, align 8
  %4 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %2, i32 0, i32 1
  store i64 4, ptr %4, align 4
  %5 = load %"github.com/goplus/llgo/internal/runtime.String", ptr %2, align 8
  call void @"github.com/goplus/llgo/internal/runtime.PrintString"(%"github.com/goplus/llgo/internal/runtime.String" %5)
  call void @"github.com/goplus/llgo/internal/runtime.PrintByte"(i8 10)
  %6 = alloca %main.bar, align 8
  %7 = call ptr @"github.com/goplus/llgo/internal/runtime.Zeroinit"(ptr %6, i64 16)
  %8 = call { %main.bar, i1 } @main.Foo(%"github.com/goplus/llgo/internal/runtime.eface" zeroinitializer)
  %9 = extractvalue { %main.bar, i1 } %8, 0
  store %main.bar %9, ptr %7, align 8
  %10 = extractvalue { %main.bar, i1 } %8, 1
  %11 = getelementptr inbounds %main.bar, ptr %7, i32 0, i32 0
  %12 = load ptr, ptr %11, align 8
  %13 = getelementptr inbounds %main.bar, ptr %7, i32 0, i32 1
  %14 = load float, ptr %13, align 4
  %15 = xor i1 %10, true
  %16 = alloca %"github.com/goplus/llgo/internal/runtime.String", align 8
  %17 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %16, i32 0, i32 0
  store ptr @3, ptr %17, align 8
  %18 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %16, i32 0, i32 1
  store i64 6, ptr %18, align 4
  %19 = load %"github.com/goplus/llgo/internal/runtime.String", ptr %16, align 8
  call void @"github.com/goplus/llgo/internal/runtime.PrintPointer"(ptr %12)
  call void @"github.com/goplus/llgo/internal/runtime.PrintByte"(i8 32)
  %20 = fpext float %14 to double
  call void @"github.com/goplus/llgo/internal/runtime.PrintFloat"(double %20)
  call void @"github.com/goplus/llgo/internal/runtime.PrintByte"(i8 32)
  call void @"github.com/goplus/llgo/internal/runtime.PrintString"(%"github.com/goplus/llgo/internal/runtime.String" %19)
  call void @"github.com/goplus/llgo/internal/runtime.PrintByte"(i8 32)
  call void @"github.com/goplus/llgo/internal/runtime.PrintBool"(i1 %15)
  call void @"github.com/goplus/llgo/internal/runtime.PrintByte"(i8 10)
  ret i32 0
}

declare void @"github.com/goplus/llgo/internal/runtime.PrintString"(%"github.com/goplus/llgo/internal/runtime.String")

declare void @"github.com/goplus/llgo/internal/runtime.PrintByte"(i8)

declare void @"github.com/goplus/llgo/internal/runtime.init"()

declare ptr @"github.com/goplus/llgo/internal/runtime.Zeroinit"(ptr, i64)

declare void @"github.com/goplus/llgo/internal/runtime.PrintPointer"(ptr)

declare void @"github.com/goplus/llgo/internal/runtime.PrintFloat"(double)

declare void @"github.com/goplus/llgo/internal/runtime.PrintBool"(i1)

define void @"main.init$abi"() {
_llgo_0:
  %0 = alloca %"github.com/goplus/llgo/internal/runtime.String", align 8
  %1 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %0, i32 0, i32 0
  store ptr @4, ptr %1, align 8
  %2 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %0, i32 0, i32 1
  store i64 2, ptr %2, align 4
  %3 = load %"github.com/goplus/llgo/internal/runtime.String", ptr %0, align 8
  %4 = load ptr, ptr @"*_llgo_byte", align 8
  %5 = alloca %"github.com/goplus/llgo/internal/runtime.String", align 8
  %6 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %5, i32 0, i32 0
  store ptr @5, ptr %6, align 8
  %7 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %5, i32 0, i32 1
  store i64 0, ptr %7, align 4
  %8 = load %"github.com/goplus/llgo/internal/runtime.String", ptr %5, align 8
  %9 = call %"github.com/goplus/llgo/internal/abi.StructField" @"github.com/goplus/llgo/internal/runtime.StructField"(%"github.com/goplus/llgo/internal/runtime.String" %3, ptr %4, i64 0, %"github.com/goplus/llgo/internal/runtime.String" %8, i1 false)
  %10 = alloca %"github.com/goplus/llgo/internal/runtime.String", align 8
  %11 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %10, i32 0, i32 0
  store ptr @6, ptr %11, align 8
  %12 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %10, i32 0, i32 1
  store i64 1, ptr %12, align 4
  %13 = load %"github.com/goplus/llgo/internal/runtime.String", ptr %10, align 8
  %14 = call ptr @"github.com/goplus/llgo/internal/runtime.Basic"(i64 13)
  %15 = alloca %"github.com/goplus/llgo/internal/runtime.String", align 8
  %16 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %15, i32 0, i32 0
  store ptr @7, ptr %16, align 8
  %17 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %15, i32 0, i32 1
  store i64 0, ptr %17, align 4
  %18 = load %"github.com/goplus/llgo/internal/runtime.String", ptr %15, align 8
  %19 = call %"github.com/goplus/llgo/internal/abi.StructField" @"github.com/goplus/llgo/internal/runtime.StructField"(%"github.com/goplus/llgo/internal/runtime.String" %13, ptr %14, i64 8, %"github.com/goplus/llgo/internal/runtime.String" %18, i1 false)
  %20 = alloca %"github.com/goplus/llgo/internal/runtime.String", align 8
  %21 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %20, i32 0, i32 0
  store ptr @8, ptr %21, align 8
  %22 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %20, i32 0, i32 1
  store i64 4, ptr %22, align 4
  %23 = load %"github.com/goplus/llgo/internal/runtime.String", ptr %20, align 8
  %24 = call ptr @"github.com/goplus/llgo/internal/runtime.AllocU"(i64 112)
  %25 = getelementptr %"github.com/goplus/llgo/internal/abi.StructField", ptr %24, i64 0
  store %"github.com/goplus/llgo/internal/abi.StructField" %9, ptr %25, align 8
  %26 = getelementptr %"github.com/goplus/llgo/internal/abi.StructField", ptr %24, i64 1
  store %"github.com/goplus/llgo/internal/abi.StructField" %19, ptr %26, align 8
  %27 = alloca %"github.com/goplus/llgo/internal/runtime.Slice", align 8
  %28 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.Slice", ptr %27, i32 0, i32 0
  store ptr %24, ptr %28, align 8
  %29 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.Slice", ptr %27, i32 0, i32 1
  store i64 2, ptr %29, align 4
  %30 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.Slice", ptr %27, i32 0, i32 2
  store i64 2, ptr %30, align 4
  %31 = load %"github.com/goplus/llgo/internal/runtime.Slice", ptr %27, align 8
  %32 = call ptr @"github.com/goplus/llgo/internal/runtime.Struct"(%"github.com/goplus/llgo/internal/runtime.String" %23, i64 16, %"github.com/goplus/llgo/internal/runtime.Slice" %31)
  %33 = alloca %"github.com/goplus/llgo/internal/runtime.String", align 8
  %34 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %33, i32 0, i32 0
  store ptr @9, ptr %34, align 8
  %35 = getelementptr inbounds %"github.com/goplus/llgo/internal/runtime.String", ptr %33, i32 0, i32 1
  store i64 8, ptr %35, align 4
  %36 = load %"github.com/goplus/llgo/internal/runtime.String", ptr %33, align 8
  %37 = call ptr @"github.com/goplus/llgo/internal/runtime.Named"(%"github.com/goplus/llgo/internal/runtime.String" %36, ptr %32)
  store ptr %37, ptr @main.bar, align 8
  ret void
}

declare ptr @"github.com/goplus/llgo/internal/runtime.Struct"(%"github.com/goplus/llgo/internal/runtime.String", i64, %"github.com/goplus/llgo/internal/runtime.Slice")

declare %"github.com/goplus/llgo/internal/abi.StructField" @"github.com/goplus/llgo/internal/runtime.StructField"(%"github.com/goplus/llgo/internal/runtime.String", ptr, i64, %"github.com/goplus/llgo/internal/runtime.String", i1)

declare ptr @"github.com/goplus/llgo/internal/runtime.Basic"(i64)

declare ptr @"github.com/goplus/llgo/internal/runtime.AllocU"(i64)

declare ptr @"github.com/goplus/llgo/internal/runtime.Named"(%"github.com/goplus/llgo/internal/runtime.String", %"github.com/goplus/llgo/internal/runtime.String", ptr, %"github.com/goplus/llgo/internal/runtime.Slice")
