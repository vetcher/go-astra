package interfaces

type InterfaceA interface {
	A()
	B(string) error
	C(a, b int) (string, error)
	D(a interface{}) (d, e string, f error)
}

// Type documentation
type CommentInterface interface {
	A() // Method A inline comment
	// Method B documentation comment.
	B()
	/*
		Multi-line documentation of C method.
	*/
	C()
	D() /*Multi-line documentation of
	D method*/
}

// Type documentation will be in comment's block of A, B, C, D, E interfaces.
type (
	// Only in A interface.
	A interface {
		A()
	}
	// Only in B interface.
	B interface {
		B()
	}
	// This comment will not be parsed.

	// C docs.
	C interface {
		C()
	}
	/*D documentation*/
	D interface {
		D()
	}
	E interface {
		A // embedding A interface
		B // embedding B interface
	}
)

type ComplexInterface interface {
	A(a interface {
		B()
		ComplexInterface
	}) interface {
		C()
		D()
	}
}
