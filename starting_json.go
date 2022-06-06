package main

const startingConfigJson = `{
	"color_theme": "default",
	"tab_stop": 4,
	"quit_times": 1,
	"empty_line_char": " "
}`

const startingSyntaxJson = `[
    {
		"filetype":  "c",
		"filematch": [".c", ".h", "cpp", ".cc"],
		"keywords": [
			"switch", "if", "while", "for", "break", "continue", "return",
			"else", "struct", "union", "typedef", "static", "enum", "class",
			"case",

			"int|", "long|", "double|", "float|", "char|", "unsigned|",
			"signed|", "void|"
        ],
		"scs": "//",
		"mcs": "/*",
		"mce": "*/",
		"flags": {
            "highlight_numbers": true,
			"highlight_strings": true,
            "highlight_booleans": true
        }
	},
    {
        "filetype":  "go",
        "filematch": [".go"],
        "keywords": [
            "break", "default", "func", "interface", "select", "case", "defer",
            "go", "map", "struct", "chan", "else", "goto", "package", "switch",
            "const", "fallthrough", "if", "range", "type", "continue", "for",
            "import", "return", "var",

            "append|", "bool|", "byte|", "cap|", "close|", "complex|",
            "complex64|", "complex128|", "error|", "uint16|", "copy|", "false|",
            "float32|", "float64|", "imag|", "int|", "int8|", "int16|",
            "uint32|", "int32|", "int64|", "iota|", "len|", "make|", "new|",
            "nil|", "panic|", "uint64|", "print|", "println|", "real|",
            "recover|", "rune|", "string|", "true|", "uint|", "uint8|",
            "uintptr|", "delete|", "error|", "float32|", "float64|"
        ],
        "scs": "//",
        "mcs": "/*",
        "mce": "*/",
        "flags": {
            "highlight_numbers": true,
            "highlight_strings": true,
            "highlight_booleans": true
        }
    },
    {
		"filetype":  "python",
		"filematch": [".py"],
		"keywords": [
			"and", "as", "assert", "break", "class", "continue", "def", "del",
			"elif", "else", "except", "exec", "finally", "for", "from", "global",
			"if", "import", "in", "is", "lambda", "not", "or", "pass", "print",
			"raise", "return", "try", "while", "with", "yield"
        ],
		"scs": "#",
		"mcs": "#",
		"mce": "#",
		"flags": {
            "highlight_numbers": true,
            "highlight_strings": true,
            "highlight_booleans": true
        }
	},
    {
        "filetype":  "java",
        "filematch": [".java"],
        "keywords": [
            "abstract", "continue", "for", "new", "switch", "assert", "default",
            "goto", "package", "synchronized", "boolean", "do", "if", "private",
            "this", "break", "double", "implements", "protected", "throw",
            "byte", "else", "import", "public", "throws", "case", "enum",
            "instanceof", "return", "transient", "catch", "extends", "int",
            "short", "try", "char", "final", "interface", "static", "void",
            
            "int|", "long|", "double|", "float|", "char|", "unsigned|",
            "signed|", "void|"
        ],
        "scs": "//",
        "mcs": "/*",
        "mce": "*/",
        "flags": {
            "highlight_numbers": true,
            "highlight_strings": true,
            "highlight_booleans": true
        }
    },
    {
        "filetype": "javascript",
        "filematch": [".js", ".ts"],
        "keywords": [
            "break", "case", "catch", "class", "const", "continue", "debugger",
            "default", "delete", "do", "else", "enum", "export", "extends",
            "false", "finally", "for", "function", "if", "implements", "import",
            "in", "instanceof", "interface", "let", "new", "null", "package",
            "private", "protected", "public", "return", "static", "super",
            "switch", "this", "throw", "true", "try", "typeof", "var", "void",
            "while", "with", "yield",

            "int|", "long|", "double|", "float|", "char|", "unsigned|",
            "signed|", "void|"
        ],
        "scs": "//",
        "mcs": "/*",
        "mce": "*/",
        "flags": {
            "highlight_numbers": true,
            "highlight_strings": true,
            "highlight_booleans": true
        }
    },
    {
        "filetype": "rust",
        "filematch": [".rs"],
        "keywords": [
            "break", "case", "catch", "class", "const", "continue", "debugger",
            "default", "delete", "do", "else", "enum", "export", "extends",
            "false", "finally", "for", "function", "if", "implements", "import",
            "in", "instanceof", "interface", "let", "new", "null", "package",
            "private", "protected", "public", "return", "static", "super",
            "switch", "this", "throw", "true", "try", "typeof", "var", "void",
            "while", "with", "yield",

            "int|", "long|", "double|", "float|", "char|", "unsigned|",
            "signed|", "void|"
        ],
        "scs": "//",
        "mcs": "/*",
        "mce": "*/",
        "flags": {
            "highlight_numbers": true,
            "highlight_strings": true,
            "highlight_booleans": true
        }
    }
]`
