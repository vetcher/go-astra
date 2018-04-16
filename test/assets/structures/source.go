package structures

type MainStructure struct {
	A struct {
		A struct {
		}
		B map[struct {
			A interface{}
			B string
		}]struct {
			A int
			B chan struct{}
			C chan<- struct{}
			D <-chan struct{}
		} `json:"b"xml:"b"`
	}
	B struct {
		A struct {
			A int
		}
		B *struct {
			A **struct {
				A []struct {
					A int // comment of A
				}
			}
		}
		C func(struct {
			A int
		})
	}
}
