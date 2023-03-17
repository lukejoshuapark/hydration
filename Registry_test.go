package hydration

import (
	"fmt"
	"testing"

	"github.com/lukejoshuapark/test"
	"github.com/lukejoshuapark/test/does"
	"github.com/lukejoshuapark/test/is"
)

type dummyType struct{}

type testCase struct {
	ExplicitTypeKeyWhenRegistering  bool
	ExplicitRegistryWhenRegistering bool
	ExplicitTypeKeyWhenResolving    bool
	ExplicitRegistryWhenResolving   bool
	ExplicitTypeKeyWhenChecking     bool
	ExplicitRegistryWhenChecking    bool
}

func createTestCase(i int) testCase {
	return testCase{
		ExplicitTypeKeyWhenRegistering:  i&1 == 1,
		ExplicitRegistryWhenRegistering: i&2 == 2,
		ExplicitTypeKeyWhenResolving:    i&4 == 4,
		ExplicitRegistryWhenResolving:   i&8 == 8,
		ExplicitTypeKeyWhenChecking:     i&16 == 16,
		ExplicitRegistryWhenChecking:    i&32 == 32,
	}
}

func (t *testCase) register() {
	if t.ExplicitRegistryWhenRegistering {
		if t.ExplicitTypeKeyWhenRegistering {
			RegisterIntoWithKey[dummyType](DefaultRegistry, "hydration.dummyType")
		} else {
			RegisterInto[dummyType](DefaultRegistry)
		}
	} else {
		if t.ExplicitTypeKeyWhenRegistering {
			RegisterWithKey[dummyType]("hydration.dummyType")
		} else {
			Register[dummyType]()
		}
	}
}

func (t *testCase) check() bool {
	if t.ExplicitRegistryWhenChecking {
		if t.ExplicitTypeKeyWhenChecking {
			return KnowsTypeKeyIn(DefaultRegistry, "hydration.dummyType")
		} else {
			return KnowsIn[dummyType](DefaultRegistry)
		}
	} else {
		if t.ExplicitTypeKeyWhenChecking {
			return KnowsTypeKey("hydration.dummyType")
		} else {
			return Knows[dummyType]()
		}
	}
}

func (t *testCase) resolve() dummyType {
	if t.ExplicitRegistryWhenResolving {
		if t.ExplicitTypeKeyWhenResolving {
			return ResolveFromWithKey[dummyType](DefaultRegistry, "hydration.dummyType")
		} else {
			return ResolveFrom[dummyType](DefaultRegistry)
		}
	} else {
		if t.ExplicitTypeKeyWhenResolving {
			return ResolveWithKey[dummyType]("hydration.dummyType")
		} else {
			return Resolve[dummyType]()
		}
	}
}

func TestHydration(t *testing.T) {
	for i := 0; i < 1<<6; i++ {
		c := createTestCase(i)

		t.Run(fmt.Sprintf("Case %v", i), func(t *testing.T) {
			c.register()

			exists := c.check()
			test.That(t, exists, is.True)
			test.That(t, func() { c.resolve() }, does.NotPanic)
		})
	}
}
