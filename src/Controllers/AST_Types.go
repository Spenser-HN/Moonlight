package syntax

type AST struct {
	Type string
	Body []Body
}

type Body struct {
	Type         string
	Declarations []string
}

type Node struct {
	Id    int
	Val   interface{}
	Left  *Node
	Right *Node
}

type ObjectPattern struct {
	Type       string // 'ObjectPattern'
	Properties []Property
}

type Property struct {
	Type      string // 'Property'
	Key       Expression
	Computed  bool
	Value     Expression // Expression | nil
	Kind      string     // 'Get' | 'Set' | 'Init'
	Method    bool       // false
	Shorthand bool
}

// ----------------------- Expressions ----------------------

type Expression = any /*

ThisExpression | Identifier | Literal |

ArrayExpression | ObjectExpression | FunctionExpression |

ArrowFunctionExpression | ClassExpression | TaggedTemplateExpression |

MemberExpression | Super | MetaProperty | NewExpression | CallExpression |

UpdateExpression | AwaitExpression | UnaryExpression |

BinaryExpression | LogicalExpression | ConditionalExpression |

YieldExpression | AssignmentExpression | SequenceExpression

*/

type BindingPattern = any //ArrayPattern | ObjectPattern;

type ArrayPattern struct {
	Type     string // "ArrayPatter"
	Elements []ArrayPatternElement
}

type ArrayPatternElement = any // AssignmentPattern | Identifier | BindingPattern | RestElement | null;

type RestElement struct {
	Type     string // "RestElement"
	Argument any    // Identifier | BindingPattern;
}
type ThisExpression struct {
	Type string // 'ThisExpression'
}

type Identifier struct {
	Type string // 'Identifier'
	Name string
}

type Literal struct {
	Type  string // 'Literal'
	Value any
	Regex Literal_Regex // nil | Literal_Regex
}

type Literal_Regex struct {
	Pattern string
	Flags   string
}

type FunctionExpression struct {
	Type       string     // 'FunctionExpression'
	Id         Identifier // Identifier | nil
	Params     []FunctionParameter
	Body       BlockStatement
	Generator  bool
	Async      bool
	Expression bool
}

type FunctionParameter = any // AssignmentPattern | Identifier | BindingPattern

type ArrowFunctionExpression struct {
	Type       string // 'ArrowFunctionExpression'
	Id         Identifier
	Params     []FunctionParameter
	Body       any // BlockStatement | Expression
	Async      bool
	Expression bool //false
}

type ClassExpression struct {
	Type       string // 'ClassExpression'
	Id         Identifier
	SuperClass Identifier
	Body       ClassBody
}

type ClassBody struct {
	Type string // 'ClassBody'
	Body []MethodDefinition
}

type MethodDefinition struct {
	Type     string // 'MethodDefinition'
	Key      Expression
	Computed bool
	Value    FunctionExpression
	Kind     string // 'Method' | 'Constructor'
	Static   bool
}

type TaggedTemplateExpression struct {
	Type  string // 'TaggedTemplateExpression'
	Tag   Expression
	Quasi TemplateLiteral
}

type TemplateElement struct {
	Type  string // 'TemplateElement'
	Value struct {
		Cooked string
		Raw    string
	}
	Tail bool
}

type TemplateLiteral struct {
	Type        string // 'TemplateLiteral'
	Quasis      []TemplateElement
	Expressions []Expression
}

type MemberExpression struct {
	Type     string // 'MemberExpression'
	Computed bool
	Object   Expression
	Property Expression
}

type Super struct {
	Type string // 'Super'
}

type MetaProperty struct {
	Type     string // 'MetaProperty'
	Meta     Identifier
	Property Identifier
}

type CallExpression struct {
	Type      string // 'CallExpression'
	Callee    Expression
	Argumnets []ArgumentListElement
}

type NewExpression struct {
	Type      string // 'NewExpression'
	Callee    Expression
	Arguments []ArgumentListElement
}

type ArgumentListElement = any // Expression | SpreadElement

type SpreadElement struct {
	Type     string // 'SpreadElement'
	Argument Expression
}

type ArrayExpression struct {
	Type     string // 'ArrayExpression'
	Elements []ArrayExpressionElement
}

type ArrayExpressionElement = any //Expression | SpreadElement

type UpdateExpression struct {
	Type     string // 'UpdateExpression'
	Operator string // '++' || '--'
	Argument Expression
	Prefix   bool
}

type AwaitExpression struct {
	Type     string // 'AwaitExpression'
	Argument Expression
}

type UnaryExpression struct {
	Type     string // 'UnaryExpression'
	Operator string // '+' | '-' | '~' | '!' | 'delete' | 'void' | 'typeof';
	Argument Expression
	Prefix   bool //true
}

type BynaryExpression struct {
	Type     string // 'BinaryExpression'
	Operator string /* 'instanceof' | 'in' | '+' | '-' | '*' | '/' | '%' |
	'**' | '|' | '^' | '&' | '==' | '!=' | '===' | '!==' |
	* '<' | '>' | '<=' | '<<' | '>>' | '>>>';
	**/
	Left  Expression
	Right Expression
}

type LogicalExpression struct {
	Type     string // 'LogicalExpression'
	Operator string // '||' | '&&'
	Left     Expression
	Right    Expression
}

type ConditionalExpression struct {
	Type       string // 'ConditionalExpression'
	Test       Expression
	Consequent Expression
	Alternate  Expression
}

type YieldExpression struct {
	Type     string     // 'YieldExpression'
	Argument Expression // Expression | nil
	Delegate bool
}

type AssingmentExpression struct {
	Type     string // 'AssingmentExpression'
	Operator string /** '=' | '*=' | '**=' | '/=' | '%=' | '+=' | '-=' |
	'<<=' | '>>=' | '>>>=' | '&=' | '^=' | '|=' **/
	Left  Expression
	Right Expression
}

type SequenceExpression struct {
	Type       string // 'SequenceExpression'
	Expression []Expression
}

type Statement = any /*
   BlockStatement | BreakStatement | ContinueStatement |
   DebuggerStatement | DoWhileStatement | EmptyStatement |
   ExpressionStatement | ForStatement | ForInStatement |
   ForOfStatement | FunctionDeclaration | IfStatement |
   LabeledStatement | ReturnStatement | SwitchStatement |
   ThrowStatement | TryStatement | VariableDeclaration |
   WhileStatement | WithStatement
*/

type Declaration = any // ClassDeclaration | FunctionDeclaration |  VariableDeclaration

type StatementListItem = any // Declaration | Statement

type BlockStatement struct {
	Type string // 'BlockStatement'
	Body []StatementListItem
}

type BreakStatement struct {
	Type  string // 'BreakStatement'
	Label Identifier
}

type ClassDeclaration struct {
	Type       string // 'ClassDeclaration'
	Id         Identifier
	SuperClass Identifier
	Body       ClassBody
}

type ContinueStatement struct {
	Type  string // 'ContinueStatement'
	Label Identifier
}

type DebuggerStatement struct {
	Type string // 'DebuggerStatement'
}

type DoWhileStatement struct {
	Type string // 'DoWhileStatement'
	Body Statement
	Test Expression
}

type EmptyStatement struct {
	Type string // 'EmptyStatement'
}

type ExpressionStatement struct {
	Type       string // 'ExpressionStatement'
	Expression Expression
	Directive  string
}

type ForStatement struct {
	Type   string // 'ForStatement'
	Init   any    // Expression | VariableDeclaration
	Test   Expression
	Update Expression
	Body   Statement
}

type ForInStatement struct {
	Type  string // 'ForInStatement'
	Left  Expression
	Right Expression
	Body  Statement
	Each  bool //false
}

type ForOfStatement struct {
	Type  string // 'ForOfStatement'
	Left  Expression
	Right Expression
	Body  Statement
}

type FunctionDeclaration struct {
	Type       string // 'FunctionDeclaration'
	Id         Identifier
	Params     []FunctionParameter
	Body       BlockStatement
	Generator  bool
	Async      bool
	Expression bool //false
}

type LabeledStatement struct {
	Type  string // 'LabeledStatement'
	Label Identifier
	Body  Statement
}

type ReturnStatement struct {
	Type     string // 'ReturnStatement'
	Argument Expression
}

type SwitchStatement struct {
	Type         string // 'SwitchStatement'
	Discriminant Expression
	Cases        []SwitchCase
}

type SwitchCase struct {
	Type       string // 'SwitchCase'
	Test       Expression
	Consequent []Statement
}

type ThrowStatement struct {
	Type     string // 'ThrowStatement'
	Argument Expression
}

type TryStatement struct {
	Type      string // 'TryStatement'
	Block     BlockStatement
	Handler   CatchClause
	Finalizer BlockStatement
}

type CatchClause struct {
	Type  string // 'CatchClause'
	Param any    // Identifier | BindingPattern
	Body  BlockStatement
}

type VariableDeclaration struct {
	Type         string // "VariableDeclaration"
	Declarations []VariableDeclarator
	Kind         string // "var" | "const" | "let"
}

type VariableDeclarator struct {
	Type string     // "VariableDeclarator"
	Id   any        // Identifier | BindingPattern
	Init Expression // Expression | null
}

type WhileStatement struct {
	Type string // "WhileStatement"
	Test Expression
	Body Statement
}

type WhiteStatement struct {
	Type   string // "WhiteStatement"
	Object Expression
	Body   Statement
}

// ? -------- Script and Modules ---------- ?

type Program struct {
	Type       string // "Program"
	SourceType string // "Script" | "Module"
	Body       []any  // StatementListItem | ModuleItem
}

type ModuleItem = any //ImportDeclaration | ExportDeclaration | StatementListItem

type ImportDeclaration struct {
	Type       string // "ImportDeclaration"
	Specifiers []ImportSpecifiers
	Source     Literal
}

type ImportSpecifiers struct {
	Type     string // 'ImportSpecifier' | 'ImportDefaultSpecifier' | 'ImportNamespaceSpecifier'
	Local    Identifier
	Imported Identifier
}

type ExportDeclaration = any // ExportAllDeclaration | ExportDefaultDeclaration | ExportNamedDeclaration

type ExportAllDeclaration struct {
	Type   string // "ExportAllDeclaration"
	Source Literal
}

type ExportDefaultDeclaration struct {
	Type        string // "ExportDefaultDeclaration"
	Declaration any    //  Identifier | BindingPattern | ClassDeclaration | Expression | FunctionDeclaration
}

type ExportNamedDeclaration struct {
	Type        string // "ExportNamedDeclaration"
	Declaration any    // ClassDeclaration | FunctionDeclaration | VariableDeclaration;
	Specifiers  []ExportSpecifier
	Source      Literal
}

type ExportSpecifier struct {
	Type     string // "ExportSpecifier"
	Exported Identifier
	Imported Identifier
}
