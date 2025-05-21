// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package redblack

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str/nodes"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"testing"
)

// TestDeleteFixup tests the fixup after a deletion
//
// after a fixup we must assert
// 1. the node can't be found on the tree
// 2. the tree remains balanced
func TestDeleteFixup(t *testing.T) {
	cmp := comparators.Ordered[int]{}
	t.Run("Balanced black-red siblings", func(t *testing.T) {
		// Balanced tree:
		//      a(10)(Root)
		//        /    \
		//  b(5)[B]    c(15)[R]
		//             /    \
		//      d(12)[B]    e(20)[B]
		a := Root(10)
		b := Black(5, a)
		a.left = b
		c := Red(15, a)
		a.right = c
		d := Black(12, c)
		c.left = d
		e := Black(20, c)
		c.right = e
		root := a
		var parent *Node[int]
		// Test 1: transplant b(5)
		// a) assert the node can be found
		found, ok := FindNode(root, b.value, cmp.Compare)
		if !ok {
			t.Errorf("Node %v not found on the tree", b.value)
		}
		if found != b {
			t.Errorf("Expected node to be %v, but got %v", b, found)
		}
		// b) remove node
		root, parent, _ = halfTransplantLeft(root, b)
		// After deletion (not balanced)
		//      a(10)(Root)
		//              \
		//  	       c(15)[R]
		//             /    \
		//      d(12)[B]    e(20)[B]
		// Perform fixup (root, deletedParent)
		root = deleteFixup(root, parent, true)
		// Expected result (balanced):
		//      c(15)(Root)
		//        /    \
		//  a(10)[B]   e(20)[B]
		//       \
		//     d(12)[R]
		assertValidRoot(t, "c", c, root)
		if root.left != a {
			t.Errorf("expected left child of root to be a(%d), but got %v", a.value, root.left.value)
		}
		if a.parent != root {
			t.Errorf("expected parent of a to be root, but got %v", a.parent)
		}
		if a.color != nodes.Black {
			t.Errorf("expected left child of root to be black, but got %s", a.color)
		}
		if root.right != e {
			t.Errorf("expected right child of root to be e(%d), but got %v", e.value, root.right.value)
		}
		if e.parent != root {
			t.Errorf("expected parent of e to be root, but got %v", e.parent)
		}
		if e.color != nodes.Black {
			t.Errorf("expected right child of root to be red, but got %s", root.right.color)
		}
		if a.left != nil {
			t.Errorf("expected left child of a to be nil, but got %v", a.left)
		}
		if a.right != d {
			t.Errorf("expected right child of a to be d(%d), but got %v", d.value, a.right.value)
		}
		if e.left != nil {
			t.Errorf("expected left child of e to be nil, but got %v", e.left)
		}
		if e.right != nil {
			t.Errorf("expected right child of e to be nil, but got %v", e.right)
		}
		if d.parent != a {
			t.Errorf("expected parent of d to be a, but got %v", d.parent)
		}
		if d.color != nodes.Red {
			t.Errorf("expected d color to be Black, but got %s", d.color)
		}
		if d.left != nil {
			t.Errorf("expected left child of d to be nil, but got %v", d.left)
		}
		if d.right != nil {
			t.Errorf("expected right child of d to be nil, but got %v", d.right)
		}

		// c) assert node can't be found
		_, ok = FindNode(root, b.value, cmp.Compare)
		if ok {
			t.Errorf("Node %v should not be found on the tree", b.value)
		}

		// Test 2: delete e(20) (causes unbalance)
		// From:
		//      c(15)(Root)
		//        /    \
		//  a(10)[B]   e(20)[B]
		//       \
		//     d(12)[R]

		// a) assert the node can be found
		found, ok = FindNode(root, e.value, cmp.Compare)
		if !ok {
			t.Errorf("Node e(%d) not found on the tree", e.value)
		}
		if found != e {
			t.Errorf("Expected node to be 'e(%d)', but got %v", e.value, found)
		}

		// b) remove node
		root, parent, _ = halfTransplantRight(root, e)
		// To (unbalanced):
		//      c(15)(Root)
		//        /
		//  a(10)[B]
		//       \
		//     d(12)[R]
		// Perform fixup
		root = deleteFixup(root, parent, false)
		// Expected result (balanced):
		//         d(12)(Root)
		//          /    \
		//    a(10)[B]  c(15)[B]
		assertValidRoot(t, "d", d, root)
		if root.left != a {
			t.Errorf("expected left child of root to be a(%d), but got %v", a.value, root.left.value)
		}
		if a.parent != root {
			t.Errorf("expected parent of a to be root, but got %v", a.parent)
		}
		if a.color != nodes.Black {
			t.Errorf("expected left child of root to be black, but got %s", a.color)
		}
		if root.right != c {
			t.Errorf("expected right child of root to be c(%d), but got %v", c.value, root.right.value)
		}
		if c.parent != root {
			t.Errorf("expected parent of c to be root, but got %v", c.parent)
		}
		if c.color != nodes.Black {
			t.Errorf("expected right child of root to be red, but got %s", root.right.color)
		}
		if a.left != nil {
			t.Errorf("expected left child of a to be nil, but got %v", a.left)
		}
		if a.right != nil {
			t.Errorf("expected right child of a to be nil, but got %v", a.right)
		}
		if c.left != nil {
			t.Errorf("expected left child of c to be nil, but got %v", c.left)
		}
		if c.right != nil {
			t.Errorf("expected right child of c to be nil, but got %v", c.right)
		}

		// c) assert node can't be found
		_, ok = FindNode(root, e.value, cmp.Compare)
		if ok {
			t.Errorf("Node 'e(%d)' should not be found on the tree", e.value)
		}
	})

	t.Run("Occasional Imbalance of left", func(t *testing.T) {
		// Create a valid Red-Black Tree:
		//      a(10)[Root]
		//        /    \
		//  b(5)[B]    c(15)[R]
		//                   \
		//                  d(20)[B]
		a := Root(10)
		a.left = Black(5, a)
		b := a.left
		a.right = Red(15, a)
		c := a.right
		c.right = Black(20, c)
		d := c.right
		root := a
		var parent *Node[int]
		// Delete node b (causes color imbalance)
		root, parent, _ = halfTransplantLeft(root, b)
		// From:
		//      a(10)[Root]
		//        /    \
		//      nil    c(15)[R]
		//                   \
		//                  d(20)[B]
		// Perform fixup
		root = deleteFixup(root, parent, true)
		// To: Expected structure:
		//      c(15)(Root)
		//        /    \
		// a(10)[B]    d(20)[B]
		assertValidRoot(t, "c", c, root)
		if root.left != a {
			t.Errorf("Expected left child of root to be a(%d), but got %v", a.value, root.left.value)
		}
		if a.parent != root {
			t.Errorf("Expected parent of a to be root, but got %v", a.parent)
		}
		if a.color != nodes.Black {
			t.Errorf("Expected left child of root to be Red, but got %s", root.left.color)
		}
		if a.left != nil {
			t.Errorf("Expected left child of a to be nil, but got %v", a.left)
		}
		if a.right != nil {
			t.Errorf("Expected right child of a to be nil, but got %v", a.right)
		}
		if root.right != d {
			t.Errorf("Expected right child of root to be d(%d), but got %v", d.value, root.right.value)
		}
		if d.parent != root {
			t.Errorf("Expected parent of d to be root, but got %v", d.parent)
		}
		if d.color != nodes.Black {
			t.Errorf("Expected d color to be Black, but got %s", d.color)
		}
		if d.left != nil {
			t.Errorf("Expected left child of d to be nil, but got %v", d.left)
		}
		if d.right != nil {
			t.Errorf("Expected right child of d to be nil, but got %v", d.right)
		}
	})

	t.Run("Occasional imbalance on right", func(t *testing.T) {
		// Create a tree:
		//      a(10)[Root]
		//        /    \
		//   b(5)[R]  c(15)[B]
		//     /
		// d(3)[B]
		a := Root(10)
		a.left = Red(5, a)
		b := a.left
		a.right = Black(15, a)
		c := a.right
		b.left = Black(3, b)
		d := b.left
		root := a
		var parent *Node[int]
		// Delete node c(15) (causes color imbalance)
		root, parent, _ = halfTransplantRight(root, c)
		// From (color imbalance):
		//      a(10)[Root]
		//         /
		//     b(5)[R]
		//      /
		//  d(3)[B]
		// Perform fixup
		root = deleteFixup(root, parent, false)
		// Expected structure:
		//       b(5)[Root]
		//        /    \
		//   d(3)[B]  a(10)[B]
		assertValidRoot(t, "b", b, root)
		if b.left != d {
			t.Errorf("Expected left child of b to be d(%d), but got %v", d.value, b.left.value)
		}
		if d.parent != b {
			t.Errorf("Expected parent of d to be b, but got %v", d.parent)
		}
		if d.color != nodes.Black {
			t.Errorf("Expected d color to be Black, but got %s", d.color)
		}
		if d.left != nil {
			t.Errorf("Expected left child of d to be nil, but got %v", d.left)
		}
		if d.right != nil {
			t.Errorf("Expected right child of d to be nil, but got %v", d.right)
		}
		if b.right != a {
			t.Errorf("Expected right child of b to be a(%d), but got %v", a.value, b.right.value)
		}
		if a.parent != b {
			t.Errorf("Expected parent of a to be b, but got %v", a.parent)
		}
		if a.color != nodes.Black {
			t.Errorf("Expected a color to be Black, but got %s", a.color)
		}
		if a.left != nil {
			t.Errorf("Expected left child of a to be nil, but got %v", a.left)
		}
		if a.right != nil {
			t.Errorf("Expected right child of a to be nil, but got %v", a.right)
		}
	})

	t.Run("Larger Tree", func(t *testing.T) {
		// Create a valid Red-Black Tree:
		//                        a(30)[Root]
		//                       /          \
		//               b(20)[B]           c(40)[B]
		//              /      \            /      \
		//      d(15)[R]       e(25)[R]  f(35)[R]   g(50)[R]
		//        / \            /             \         \
		//h(10)[B]  i(18)[B] j(23)[B]      k(38)[B]      l(55)[B]
		a := Root(30)
		b := Black(20, a)
		c := Black(40, a)
		d := Red(15, b)
		e := Red(25, b)
		f := Red(35, c)
		g := Red(50, c)
		h := Black(10, d)
		i := Black(18, d)
		j := Black(23, e)
		k := Black(38, f)
		l := Black(55, g)

		a.left = b
		a.right = c
		b.left = d
		b.right = e
		c.left = f
		c.right = g
		d.left = h
		d.right = i
		e.left = j
		f.right = k
		g.right = l

		root := a
		var parent *Node[int]
		// Case 1: Delete node l(55) (causes unbalance)
		root, parent, _ = halfTransplantRight(root, l)
		// From:
		//                        a(30)[Root]
		//                       /          \
		//               b(20)[B]           c(40)[B]
		//              /      \            /      \
		//      d(15)[R]       e(25)[R]  f(35)[R]   g(50)[R]
		//        / \            /             \         \
		//h(10)[B]  i(18)[B] j(23)[B]      k(38)[B]       x
		root = deleteFixup(root, parent, false)
		// Expected structure after fix up:
		//                        a(30)[Root]
		//                       /          \
		//               b(20)[B]           c(40)[B]
		//              /      \            /      \
		//      d(15)[R]       e(25)[R]  f(35)[R]   g(50)[B]
		//        / \            /             \
		//h(10)[B]  i(18)[B] j(23)[B]      k(38)[B]
		assertValidRoot(t, "a", a, root)

		if root.left != b {
			t.Errorf("Expected left child of a to be b(%d), but got %v", b.value, root.left.value)
		}
		if b.parent != root {
			t.Errorf("Expected parent of b to be a, but got %v", b.parent)
		}
		if b.color != nodes.Black {
			t.Errorf("Expected b color to be Black, but got %s", b.color)
		}
		if b.left != d {
			t.Errorf("Expected left child of b to be d(%d), but got %v", d.value, b.left.value)
		}
		if b.right != e {
			t.Errorf("Expected right child of b to be e(%d), but got %v", e.value, b.right.value)
		}

		if root.right != c {
			t.Errorf("Expected right child of a to be c(%d), but got %v", c.value, root.right.value)
		}
		if c.parent != root {
			t.Errorf("Expected parent of c to be a, but got %v", c.parent)
		}
		if c.color != nodes.Black {
			t.Errorf("Expected c color to be Black, but got %s", c.color)
		}
		if c.left != f {
			t.Errorf("Expected left child of c to be f(%d), but got %v", f.value, c.left.value)
		}
		if c.right != g {
			t.Errorf("Expected right child of c to be g(%d), but got %v", g.value, c.right.value)
		}

		if d.parent != b {
			t.Errorf("Expected parent of d to be b, but got %v", d.parent)
		}
		if d.color != nodes.Red {
			t.Errorf("Expected d color to be Red, but got %s", d.color)
		}
		if d.left != h {
			t.Errorf("expected left child of d to be h(%d), but got %v", h.value, d.left.value)
		}
		if d.right != i {
			t.Errorf("expected right child of d to be i(%d), but got %v", i.value, d.right.value)
		}

		if e.parent != b {
			t.Errorf("Expected parent of e to be b, but got %v", e.parent)
		}
		if e.color != nodes.Red {
			t.Errorf("Expected e color to be Red, but got %s", e.color)
		}
		if e.left != j {
			t.Errorf("Expected left child of e to be j(%d), but got %v", j.value, e.left.value)
		}
		if e.right != nil {
			t.Errorf("Expected right child of e to be nil, but got %v", e.right)
		}

		if f.parent != c {
			t.Errorf("Expected parent of f to be c, but got %v", f.parent)
		}
		if f.color != nodes.Red {
			t.Errorf("Expected f color to be Black, but got %s", f.color)
		}
		if f.left != nil {
			t.Errorf("Expected left child of f to be nil, but got %v", f.left)
		}
		if f.right != k {
			t.Errorf("Expected right child of f to be k(%d), but got %v", k.value, f.right.value)
		}

		if g.parent != c {
			t.Errorf("Expected parent of g to be c, but got %v", g.parent)
		}
		if g.color != nodes.Black {
			t.Errorf("Expected g color to be Black, but got %s", g.color)
		}
		if g.left != nil {
			t.Errorf("Expected left child of g to be nil, but got %v", g.left)
		}
		if g.right != nil {
			t.Errorf("Expected right child of g nil, but got %v", g.right.value)
		}

		// Case 2: Delete g(50)[B] causes unbalance
		root, parent, _ = halfTransplantRight(root, g)
		//                        a(30)[Root]
		//                       /          \
		//               b(20)[B]           c(40)[B]
		//              /      \            /      \
		//      d(15)[R]       e(25)[R]  f(35)[R]   x  <-- g(50)[B]
		//        / \            /             \
		//h(10)[B]  i(18)[B] j(23)[B]      k(38)[B]
		root = deleteFixup(root, parent, false)
		// Expected structure:
		//                        a(30)[Root]
		//                       /          \
		//               b(20)[B]           f(35)[B]
		//              /      \                 \
		//      d(15)[R]       e(25)[R]         c(40)[B]
		//        / \            /             /
		//h(10)[B]  i(18)[B] j(23)[B]      k(38)[R]
		assertValidRoot(t, "a", root, a)
		if b.color != nodes.Black {
			t.Errorf("expected d color to be Black, but got %s", b.color)
		}
		if b.parent != a {
			t.Errorf("Expected parent of b to be a(%d), but got %v", a.value, b.parent)
		}
		if b.left != d {
			t.Errorf("expected left child of b to be d(%d), but got %v", d.value, b.left.value)
		}
		if b.right != e {
			t.Errorf("expected right child of b to be e(%d), but got %v", e.value, b.right.value)
		}

		if f.color != nodes.Black {
			t.Errorf("expected f color to be black but got %s", c.color)
		}
		if f.parent != a {
			t.Errorf("expected f parent to be a(%d), but got %v", a.value, f.parent.value)
		}
		if f.left != nil {
			t.Errorf("expected f left child to be nil but got %v", f.left)
		}
		if f.right != c {
			t.Errorf("expected f right child to be c(%d), but got %v", c.value, f.right.value)
		}

		if d.color != nodes.Red {
			t.Errorf("expected d color to be red but got %s", d.color)
		}
		if d.parent != b {
			t.Errorf("expected d parent to be b(%d), but got %v", b.value, d.parent.value)
		}
		if d.left != h {
			t.Errorf("expected d left child to be h(%d), but got %v", h.value, d.left.value)
		}
		if d.right != i {
			t.Errorf("expected d right child to be i(%d), but got %v", i.value, d.right.value)
		}

		if e.color != nodes.Red {
			t.Errorf("expected e color to be red but got %s", e.color)
		}
		if e.parent != b {
			t.Errorf("expected e parent to be b(%d), but got %v", b.value, e.parent.value)
		}
		if e.left != j {
			t.Errorf("expected e left child to be j(%d), but got %v", j.value, e.left.value)
		}
		if e.right != nil {
			t.Errorf("expected e right child to be nil, but got %v", e.right.value)
		}

		if c.color != nodes.Black {
			t.Errorf("expected c color to be black but got %s", c.color)
		}
		if c.parent != f {
			t.Errorf("expected c parent to be f(%d), but got %v", f.value, c.parent.value)
		}
		if c.left != k {
			t.Errorf("expected c left child to be k(%d), but got %v", k.value, c.left.value)
		}
		if c.right != nil {
			t.Errorf("expected c.right to be nil, but got %v", c.right.value)
		}

		if h.color != nodes.Black {
			t.Errorf("expected h color to be Black, but got %s", h.color)
		}
		if h.parent != d {
			t.Errorf("Expected parent of h to be d(%d), but got %v", d.value, h.parent)
		}
		if h.left != nil {
			t.Errorf("expected h left child to be nil, but got %v", h.left.value)
		}
		if h.right != nil {
			t.Errorf("expected h right child to be nil, but got %v", h.right.value)
		}

		if i.color != nodes.Black {
			t.Errorf("expected i color to be Black, but got %s", i.color)
		}
		if i.parent != d {
			t.Errorf("Expected parent of i to be d(%d), but got %v", d.value, i.parent)
		}
		if i.left != nil {
			t.Errorf("expected i left child to be nil, but got %v", i.left.value)
		}
		if i.right != nil {
			t.Errorf("expected i right child to be nil, but got %v", i.right.value)
		}

		if j.color != nodes.Black {
			t.Errorf("expected j color to be Black, but got %s", j.color)
		}
		if j.parent != e {
			t.Errorf("Expected parent of j to be e(%d), but got %v", e.value, j.parent)
		}
		if j.left != nil {
			t.Errorf("expected j left child to be nil, but got %v", j.left.value)
		}
		if j.right != nil {
			t.Errorf("expected j right child to be nil, but got %v", j.right.value)
		}

		if k.color != nodes.Red {
			t.Errorf("expected k color to be Black, but got %s", k.color)
		}
		if k.parent != c {
			t.Errorf("Expected parent of k to be c(%d), but got %v", c.value, k.parent)
		}
		if k.left != nil {
			t.Errorf("expected k left child to be nil, but got %v", k.left.value)
		}
		if k.right != nil {
			t.Errorf("expected k right child to be nil, but got %v", k.right.value)
		}

		// Case 3: Delete c(40)[B] causes unbalance, but it doesn't fixup
		// because the replacement is red. So we delete K(38)[B]
		root, _, _ = halfTransplantLeft(root, c)
		f.right.color = nodes.Black // color fixup
		root, parent, _ = halfTransplantRight(root, k)
		//                        a(30)[Root]
		//                       /          \
		//               b(20)[B]           f(35)[B]
		//              /      \                 \
		//      d(15)[R]       e(25)[R]          x   <-- K(38)[B]
		//        / \            /
		//h(10)[B]  i(18)[B] j(23)[B]
		root = deleteFixup(root, parent, false)
		// Expected Tree:
		//              b(20)[Root]
		//           /              \
		//     d(15)[B]          a(30)[B]
		//      /    \              /   \
		// h(10)[B] i(18)[B]   e(25)[R] f(35)[B]
		//                       /
		//                  j(23)[B]
		assertValidRoot(t, "b", b, root)
		if d.color != nodes.Black {
			t.Errorf("expected d to be black but got '%s'", d.color)
		}
		if d.parent != b {
			t.Errorf("Expected parent of b to be a(%d), but got %v", a.value, b.parent)
		}
		if d.left != h {
			t.Errorf("expected left child of d to be h(%d) but got %v", h.value, d.left.value)
		}
		if d.right != i {
			t.Errorf("expected right child of d to be i(%d) but got %v", i.value, d.right.value)
		}

		if a.color != nodes.Black {
			t.Errorf("expected a to be black but got '%s'", a.color)
		}
		if a.parent != b {
			t.Errorf("Expected parent of a to be b(%d), but got %v", b.value, a.parent)
		}
		if a.left != e {
			t.Errorf("expected a left child to be e(%d) but got %v", e.value, a.left.value)
		}
		if a.right != f {
			t.Errorf("expect a right child to be f(%d) but got %v", f.value, a.right.value)
		}
		if h.color != nodes.Black {
			t.Errorf("expected h to be black but got '%s'", h.color)
		}
		if h.parent != d {
			t.Errorf("Expected parent of h to be d(%d), but got %v", d.value, h.parent)
		}
		if h.left != nil {
			t.Errorf("expected h left child to be nil, but got %v", h.left.value)
		}
		if h.right != nil {
			t.Errorf("expected h right child to be nil, but got %v", h.right.value)
		}

		if i.color != nodes.Black {
			t.Errorf("expected i to be black but got '%s'", i.color)
		}
		if i.parent != d {
			t.Errorf("Expected parent of i to be d(%d), but got %v", d.value, i.parent)
		}
		if i.left != nil {
			t.Errorf("expected i left child to be nil, but got %v", i.left.value)
		}
		if i.right != nil {
			t.Errorf("expected i right child to be nil, but got %v", i.right.value)
		}

		if e.color != nodes.Red {
			t.Errorf("expected e to be red but got '%s'", e.color)
		}
		if e.parent != a {
			t.Errorf("Expected parent of e to be a(%d), but got %v", a.value, e.parent)
		}
		if e.left != j {
			t.Errorf("expected left child of e to be j(%d) but got %v", j.value, e.left.value)
		}
		if e.right != nil {
			t.Errorf("expected right child of e to be nil, but got %v", e.right.value)
		}

		if f.color != nodes.Black {
			t.Errorf("expected f color to be black but got '%s'", f.color)
		}
		if f.parent != a {
			t.Errorf("expected f parent to be a(%d) but got %v", a.value, f.parent)
		}
		if f.left != nil {
			t.Errorf("expected a left child to be nil, but got %v", a.left.value)
		}
		if f.right != nil {
			t.Errorf("expected a right child to be nil, but got %v", a.right.value)
		}
	})
}

// TestDeleteNode tests a node deletion
//
// after the deletion we must assert:
//  1. the node doesn't exist and can't be found on the tree
//  2. the tree remains balanced
func TestDeleteNode(t *testing.T) {
	// Create a valid Red-Black Tree:
	//                          a(30)[Root]
	//                         /          \
	//                 b(20)[B]           c(40)[B]
	//                /      \            /      \
	//        d(15)[R]       e(25)[R]  f(35)[R]   g(50)[R]
	//          / \            /             \         \
	//  h(10)[B]  i(18)[B] j(23)[B]      k(38)[B]      l(55)[B]
	a := Root(30)
	b := Black(20, a)
	c := Black(40, a)
	d := Red(15, b)
	e := Red(25, b)
	f := Red(35, c)
	g := Red(50, c)
	h := Black(10, d)
	i := Black(18, d)
	j := Black(23, e)
	k := Black(38, f)
	l := Black(55, g)

	a.left = b
	a.right = c
	b.left = d
	b.right = e
	c.left = f
	c.right = g
	d.left = h
	d.right = i
	e.left = j
	f.right = k
	g.right = l

	root := a

	// Case 1: Delete node l(55) (causes unbalance)
	//                        a(30)[Root]
	//                       /          \
	//               b(20)[B]           c(40)[B]
	//              /      \            /      \
	//      d(15)[R]       e(25)[R]  f(35)[R]   g(50)[R]
	//        / \            /             \         \
	//h(10)[B]  i(18)[B] j(23)[B]      k(38)[B]       x  <-- l(55)[B]
	t.Run("Delete a black leaf", func(t *testing.T) {
		root = DeleteNode(root, l)
		// Expected structure after deletion:
		//                        a(30)[Root]
		//                       /          \
		//               b(20)[B]           c(40)[B]
		//              /      \            /      \
		//      d(15)[R]       e(25)[R]  f(35)[R]   g(50)[B] <-- black
		//        / \            /             \
		//h(10)[B]  i(18)[B] j(23)[B]      k(38)[B]
		assertValidRoot(t, "a", a, root)

		if root.left != b {
			t.Errorf("Expected left child of a to be b(%d), but got %v", b.value, root.left.value)
		}
		if b.parent != root {
			t.Errorf("Expected parent of b to be a, but got %v", b.parent)
		}
		if b.color != nodes.Black {
			t.Errorf("Expected b color to be Black, but got %s", b.color)
		}
		if b.left != d {
			t.Errorf("Expected left child of b to be d(%d), but got %v", d.value, b.left.value)
		}
		if b.right != e {
			t.Errorf("Expected right child of b to be e(%d), but got %v", e.value, b.right.value)
		}

		if root.right != c {
			t.Errorf("Expected right child of a to be c(%d), but got %v", c.value, root.right.value)
		}
		if c.parent != root {
			t.Errorf("Expected parent of c to be a, but got %v", c.parent)
		}
		if c.color != nodes.Black {
			t.Errorf("Expected c color to be Black, but got %s", c.color)
		}
		if c.left != f {
			t.Errorf("Expected left child of c to be f(%d), but got %v", f.value, c.left.value)
		}
		if c.right != g {
			t.Errorf("Expected right child of c to be g(%d), but got %v", g.value, c.right.value)
		}

		if d.parent != b {
			t.Errorf("Expected parent of d to be b, but got %v", d.parent)
		}
		if d.color != nodes.Red {
			t.Errorf("Expected d color to be Red, but got %s", d.color)
		}
		if d.left != h {
			t.Errorf("expected left child of d to be h(%d), but got %v", h.value, d.left.value)
		}
		if d.right != i {
			t.Errorf("expected right child of d to be i(%d), but got %v", i.value, d.right.value)
		}

		if e.parent != b {
			t.Errorf("Expected parent of e to be b, but got %v", e.parent)
		}
		if e.color != nodes.Red {
			t.Errorf("Expected e color to be Red, but got %s", e.color)
		}
		if e.left != j {
			t.Errorf("Expected left child of e to be j(%d), but got %v", j.value, e.left.value)
		}
		if e.right != nil {
			t.Errorf("Expected right child of e to be nil, but got %v", e.right)
		}

		if f.parent != c {
			t.Errorf("Expected parent of f to be c, but got %v", f.parent)
		}
		if f.color != nodes.Red {
			t.Errorf("Expected f color to be Black, but got %s", f.color)
		}
		if f.left != nil {
			t.Errorf("Expected left child of f to be nil, but got %v", f.left)
		}
		if f.right != k {
			t.Errorf("Expected right child of f to be k(%d), but got %v", k.value, f.right.value)
		}

		if g.parent != c {
			t.Errorf("Expected parent of g to be c, but got %v", g.parent)
		}
		if g.color != nodes.Black {
			t.Errorf("Expected g color to be Black, but got %s", g.color)
		}
		if g.left != nil {
			t.Errorf("Expected left child of g to be nil, but got %v", g.left)
		}
		if g.right != nil {
			t.Errorf("Expected right child of g nil, but got %v", g.right.value)
		}
	})

	// Case 2: Delete g(50)[B] causes unbalance
	//                        a(30)[Root]
	//                       /          \
	//               b(20)[B]           c(40)[B]
	//              /      \            /      \
	//      d(15)[R]       e(25)[R]  f(35)[R]   x  <-- g(50)[B]
	//        / \            /             \
	//h(10)[B]  i(18)[B] j(23)[B]      k(38)[B]
	t.Run("Delete black node with red sibling", func(t *testing.T) {
		// 1. first ensure the previous test case state:
		if g.right != nil {
			root = DeleteNode(root, g.right)
		}
		// 2. Perform the deletion
		root = DeleteNode(root, g)
		// Expected structure:
		//                        a(30)[Root]
		//                       /          \
		//               b(20)[B]           f(35)[B]
		//              /      \                 \
		//      d(15)[R]       e(25)[R]         c(40)[B]
		//        / \            /             /
		//h(10)[B]  i(18)[B] j(23)[B]      k(38)[R]
		assertValidRoot(t, "a", root, a)
		if b.color != nodes.Black {
			t.Errorf("expected d color to be Black, but got %s", b.color)
		}
		if b.parent != a {
			t.Errorf("Expected parent of b to be a(%d), but got %v", a.value, b.parent)
		}
		if b.left != d {
			t.Errorf("expected left child of b to be d(%d), but got %v", d.value, b.left.value)
		}
		if b.right != e {
			t.Errorf("expected right child of b to be e(%d), but got %v", e.value, b.right.value)
		}

		if f.color != nodes.Black {
			t.Errorf("expected f color to be black but got %s", c.color)
		}
		if f.parent != a {
			t.Errorf("expected f parent to be a(%d), but got %v", a.value, f.parent.value)
		}
		if f.left != nil {
			t.Errorf("expected f left child to be nil but got %v", f.left)
		}
		if f.right != c {
			t.Errorf("expected f right child to be c(%d), but got %v", c.value, f.right.value)
		}

		if d.color != nodes.Red {
			t.Errorf("expected d color to be red but got %s", d.color)
		}
		if d.parent != b {
			t.Errorf("expected d parent to be b(%d), but got %v", b.value, d.parent.value)
		}
		if d.left != h {
			t.Errorf("expected d left child to be h(%d), but got %v", h.value, d.left.value)
		}
		if d.right != i {
			t.Errorf("expected d right child to be i(%d), but got %v", i.value, d.right.value)
		}

		if e.color != nodes.Red {
			t.Errorf("expected e color to be red but got %s", e.color)
		}
		if e.parent != b {
			t.Errorf("expected e parent to be b(%d), but got %v", b.value, e.parent.value)
		}
		if e.left != j {
			t.Errorf("expected e left child to be j(%d), but got %v", j.value, e.left.value)
		}
		if e.right != nil {
			t.Errorf("expected e right child to be nil, but got %v", e.right.value)
		}

		if c.color != nodes.Black {
			t.Errorf("expected c color to be black but got %s", c.color)
		}
		if c.parent != f {
			t.Errorf("expected c parent to be f(%d), but got %v", f.value, c.parent.value)
		}
		if c.left != k {
			t.Errorf("expected c left child to be k(%d), but got %v", k.value, c.left.value)
		}
		if c.right != nil {
			t.Errorf("expected c.right to be nil, but got %v", c.right.value)
		}

		if h.color != nodes.Black {
			t.Errorf("expected h color to be Black, but got %s", h.color)
		}
		if h.parent != d {
			t.Errorf("Expected parent of h to be d(%d), but got %v", d.value, h.parent)
		}
		if h.left != nil {
			t.Errorf("expected h left child to be nil, but got %v", h.left.value)
		}
		if h.right != nil {
			t.Errorf("expected h right child to be nil, but got %v", h.right.value)
		}

		if i.color != nodes.Black {
			t.Errorf("expected i color to be Black, but got %s", i.color)
		}
		if i.parent != d {
			t.Errorf("Expected parent of i to be d(%d), but got %v", d.value, i.parent)
		}
		if i.left != nil {
			t.Errorf("expected i left child to be nil, but got %v", i.left.value)
		}
		if i.right != nil {
			t.Errorf("expected i right child to be nil, but got %v", i.right.value)
		}

		if j.color != nodes.Black {
			t.Errorf("expected j color to be Black, but got %s", j.color)
		}
		if j.parent != e {
			t.Errorf("Expected parent of j to be e(%d), but got %v", e.value, j.parent)
		}
		if j.left != nil {
			t.Errorf("expected j left child to be nil, but got %v", j.left.value)
		}
		if j.right != nil {
			t.Errorf("expected j right child to be nil, but got %v", j.right.value)
		}

		if k.color != nodes.Red {
			t.Errorf("expected k color to be Black, but got %s", k.color)
		}
		if k.parent != c {
			t.Errorf("Expected parent of k to be c(%d), but got %v", c.value, k.parent)
		}
		if k.left != nil {
			t.Errorf("expected k left child to be nil, but got %v", k.left.value)
		}
		if k.right != nil {
			t.Errorf("expected k right child to be nil, but got %v", k.right.value)
		}
	})

	// Case 3: Delete c(40)[B] causes unbalance, but it doesn't fixup
	// because the replacement is red. So we delete K(38)[B]
	//                        a(30)[Root]
	//                       /          \
	//               b(20)[B]           f(35)[B]
	//              /      \                 \
	//      d(15)[R]       e(25)[R]          x   <-- K(38)[B]
	//        / \            /
	//h(10)[B]  i(18)[B] j(23)[B]
	t.Run("Case 3", func(t *testing.T) {
		// 1. Assert the previous states of the tree
		if c.right != nil {
			root = DeleteNode(root, g.right)
			root = DeleteNode(root, g)
		}
		// 2. Delete K(38)[B]
		root = DeleteNode(root, c)
		root = DeleteNode(root, k)
		// Expected Tree:
		//              b(20)[Root]
		//           /              \
		//     d(15)[B]          a(30)[B]
		//      /    \              /   \
		// h(10)[B] i(18)[B]   e(25)[R] f(35)[B]
		//                       /
		//                  j(23)[B]

		// 3. Assert the expected states of the tree
		assertValidRoot(t, "b", b, root)
		if d.color != nodes.Black {
			t.Errorf("expected d to be black but got '%s'", d.color)
		}
		if d.parent != b {
			t.Errorf("Expected parent of b to be a(%d), but got %v", a.value, b.parent)
		}
		if d.left != h {
			t.Errorf("expected left child of d to be h(%d) but got %v", h.value, d.left.value)
		}
		if d.right != i {
			t.Errorf("expected right child of d to be i(%d) but got %v", i.value, d.right.value)
		}

		if a.color != nodes.Black {
			t.Errorf("expected a to be black but got '%s'", a.color)
		}
		if a.parent != b {
			t.Errorf("Expected parent of a to be b(%d), but got %v", b.value, a.parent)
		}
		if a.left != e {
			t.Errorf("expected a left child to be e(%d) but got %v", e.value, a.left.value)
		}
		if a.right != f {
			t.Errorf("expect a right child to be f(%d) but got %v", f.value, a.right.value)
		}
		if h.color != nodes.Black {
			t.Errorf("expected h to be black but got '%s'", h.color)
		}
		if h.parent != d {
			t.Errorf("Expected parent of h to be d(%d), but got %v", d.value, h.parent)
		}
		if h.left != nil {
			t.Errorf("expected h left child to be nil, but got %v", h.left.value)
		}
		if h.right != nil {
			t.Errorf("expected h right child to be nil, but got %v", h.right.value)
		}

		if i.color != nodes.Black {
			t.Errorf("expected i to be black but got '%s'", i.color)
		}
		if i.parent != d {
			t.Errorf("Expected parent of i to be d(%d), but got %v", d.value, i.parent)
		}
		if i.left != nil {
			t.Errorf("expected i left child to be nil, but got %v", i.left.value)
		}
		if i.right != nil {
			t.Errorf("expected i right child to be nil, but got %v", i.right.value)
		}

		if e.color != nodes.Red {
			t.Errorf("expected e to be red but got '%s'", e.color)
		}
		if e.parent != a {
			t.Errorf("Expected parent of e to be a(%d), but got %v", a.value, e.parent)
		}
		if e.left != j {
			t.Errorf("expected left child of e to be j(%d) but got %v", j.value, e.left.value)
		}
		if e.right != nil {
			t.Errorf("expected right child of e to be nil, but got %v", e.right.value)
		}

		if f.color != nodes.Black {
			t.Errorf("expected f color to be black but got '%s'", f.color)
		}
		if f.parent != a {
			t.Errorf("expected f parent to be a(%d) but got %v", a.value, f.parent)
		}
		if f.left != nil {
			t.Errorf("expected 'f' left child to be nil, but got %v", f.left.value)
		}
		if f.right != nil {
			t.Errorf("expected 'f' right child to be nil, but got %v", f.right.value)
		}
	})

	// Case 4: Deleting j(23)[B] should cause color unbalance
	// which could be fixed up just by repainting e(25)[R] of black.
	//              b(20)[Root]
	//           /              \
	//     d(15)[B]          a(30)[B]
	//      /    \              /   \
	// h(10)[B] i(18)[B]   e(25)[R] f(35)[B]
	//                       /
	//                      x  <-- j(23)[B]
	t.Run("Removing black node without children and sibling", func(t *testing.T) {
		// Assert the initial structure matches with
		// the previous test
		if f.right != nil {
			root = DeleteNode(root, g.right)
			root = DeleteNode(root, g)
			root = DeleteNode(root, c)
			root = DeleteNode(root, k)
		}
		// 2. Delete j(23)[B]
		root = DeleteNode(root, j)
		// Expected structure:
		//              b(20)[Root]
		//           /              \
		//     d(15)[B]          a(30)[B]
		//      /    \              /   \
		// h(10)[B] i(18)[B]   e(25)[B] f(35)[B]
		assertValidRoot(t, "b", b, root)
		if b.left != d {
			t.Errorf("expected b left child to be d(%d), but got %v", d.value, b.left.value)
		}
		if b.right != a {
			t.Errorf("expected b right child to be a(%d), but got %v", a.value, b.right.value)
		}

		if d.parent != b {
			t.Errorf("Expected parent of d to be b(%d), but got %v", b.value, d.parent)
		}
		if d.color != nodes.Black {
			t.Errorf("expected d to be black but got '%s'", d.color)
		}
		if d.left != h {
			t.Errorf("expected left child of d to be h(%d) but got %v", h.value, d.left.value)
		}
		if d.right != i {
			t.Errorf("expected right child of d to be i(%d) but got %v", i.value, d.right.value)
		}

		if a.parent != b {
			t.Errorf("expected a parent to be b(%d), but got %v", b.value, a.parent)
		}
		if a.color != nodes.Black {
			t.Errorf("expected a to be black but got '%s'", a.color)
		}
		if a.left != e {
			t.Errorf("expected left child of a to be e(%d) but got %v", e.value, a.left.value)
		}
		if a.right != f {
			t.Errorf("expected right child of a to be f(%d) but got %v", f.value, a.right.value)
		}

		if h.parent != d {
			t.Errorf("Expected parent of h to be d(%d), but got %v", d.value, h.parent)
		}
		if h.color != nodes.Black {
			t.Errorf("expected h color to be black but got '%s'", h.color)
		}
		if h.left != nil {
			t.Errorf("expected h left child to be nil, but got %v", h.left)
		}
		if h.right != nil {
			t.Errorf("expected h right child to be nil, but got %v", h.right)
		}

		if i.parent != d {
			t.Errorf("Expected parent of i to be d(%d), but got %v", d.value, i.parent)
		}
		if i.color != nodes.Black {
			t.Errorf("expected i color to be black but got '%s'", i.color)
		}
		if i.left != nil {
			t.Errorf("expected i left child to be nil, but got %v", i.left)
		}
		if i.right != nil {
			t.Errorf("expected i right child to be nil, but got %v", i.right)
		}

		if e.parent != a {
			t.Errorf("Expected parent of e to be a(%d), but got %v", a.value, e.parent)
		}
		if e.color != nodes.Black {
			t.Errorf("expected e color to be black but got '%s'", e.color)
		}
		if e.left != nil {
			t.Errorf("expected e left child to be nil but got %v", e.left)
		}
		if e.right != nil {
			t.Errorf("expected e right child to be nil but got %v", e.right)
		}

		if f.parent != a {
			t.Errorf("Expected parent of f to be a(%d), but got %v", a.value, f.parent)
		}
		if f.color != nodes.Black {
			t.Errorf("expected f color to be black but got '%s'", f.color)
		}
		if f.left != nil {
			t.Errorf("expected f left child to be nil but got %v", f.left)
		}
		if f.right != nil {
			t.Errorf("expected f right child to be nil but got %v", f.right)
		}
	})

	// Case 5: Deleting i(18)[B] causes unbalance
	//              b(20)[Root]
	//           /              \
	//     d(15)[B]          a(30)[B]
	//      /    \              /   \
	// h(10)[B]   x <-i  e(25)[B] f(35)[B]
	t.Run("Deleting black node with black sibling", func(t *testing.T) {
		// 1. assert the tree inherits the last test final state
		if f.right != nil {
			root = DeleteNode(root, g.right)
			root = DeleteNode(root, g)
			root = DeleteNode(root, c)
			root = DeleteNode(root, k)
			root = DeleteNode(root, j)
		}
		// 2. Delete i(18)[B]
		root = DeleteNode(root, i)
		// expected structure:
		//            b(20)[Root]
		//           /          \
		//     d(15)[B]         a(30)[R]
		//      /                /   \
		// h(10)[R]        e(25)[B] f(35)[B]
		assertValidRoot(t, "b", b, root)
		if b.left != d {
			t.Errorf("expected b left child to be d(%d), but got %v", d.value, b.left.value)
		}
		if b.right != a {
			t.Errorf("expected b right child to be a(%d), but got %v", a.value, b.right.value)
		}

		if d.parent != b {
			t.Errorf("Expected parent of d to be b(%d), but got %v", b.value, d.parent)
		}
		if d.color != nodes.Black {
			t.Errorf("expected d to be black but got '%s'", d.color)
		}
		if d.left != h {
			t.Errorf("expected left child of d to be h(%d) but got %v", h.value, d.left.value)
		}
		if d.right != nil {
			t.Errorf("expected right child of d to be nil, but got %v", d.right)
		}

		if a.parent != b {
			t.Errorf("expected a parent to be b(%d), but got %v", b.value, a.parent)
		}
		if a.color != nodes.Red {
			t.Errorf("expected a to be black but got '%s'", a.color)
		}
		if a.left != e {
			t.Errorf("expected left child of a to be e(%d) but got %v", e.value, a.left.value)
		}
		if a.right != f {
			t.Errorf("expected right child of a to be f(%d) but got %v", f.value, a.right.value)
		}

		if h.parent != d {
			t.Errorf("Expected parent of h to be d(%d), but got %v", d.value, h.parent)
		}
		if h.color != nodes.Red {
			t.Errorf("expected h color to be black but got '%s'", h.color)
		}
		if h.left != nil {
			t.Errorf("expected h left child to be nil, but got %v", h.left)
		}
		if h.right != nil {
			t.Errorf("expected h right child to be nil, but got %v", h.right)
		}

		if e.parent != a {
			t.Errorf("Expected parent of e to be a(%d), but got %v", a.value, e.parent)
		}
		if e.color != nodes.Black {
			t.Errorf("expected e color to be black but got '%s'", e.color)
		}
		if e.left != nil {
			t.Errorf("expected e left child to be nil but got %v", e.left)
		}
		if e.right != nil {
			t.Errorf("expected e right child to be nil but got %v", e.right)
		}

		if f.parent != a {
			t.Errorf("Expected parent of f to be a(%d), but got %v", a.value, f.parent)
		}
		if f.color != nodes.Black {
			t.Errorf("expected f color to be black but got '%s'", f.color)
		}
		if f.left != nil {
			t.Errorf("expected f left child to be nil but got %v", f.left)
		}
		if f.right != nil {
			t.Errorf("expected f right child to be nil but got %v", f.right)
		}
	})

	// Case 6: Delete root (b(20))
	//            b(20)[Root]
	//           /          \
	//     d(15)[B]         a(30)[R]
	//      /                /   \
	// h(10)[R]        e(25)[B] f(35)[B]
	t.Run("Case 6: Delete root", func(t *testing.T) {
		// 1. assert the tree state is the same as the final state from case 5
		if f.right != nil {
			root = DeleteNode(root, g.right)
			root = DeleteNode(root, g)
			root = DeleteNode(root, c)
			root = DeleteNode(root, k)
			root = DeleteNode(root, j)
			root = DeleteNode(root, i)
		}
		// 2. delete the root
		root = DeleteNode(root, b)
		// expected final state
		//               e(25)[B]
		//             /        \
		//      d(15)[B]        a(30)[B]
		//       /                   \
		// h(10)[R]                  f(35)[R]
		assertValidRoot(t, "e", e, root)
		if e.left != d {
			t.Errorf("expected left child of 'e' to be 'd(%d)' but got '%v'", d.value, e.left.value)
		}
		if e.right != a {
			t.Errorf("expected right child of 'e' to be 'a(%d)' but got '%v'", a.value, e.right.value)
		}

		if d.parent != e {
			t.Errorf("Expected parent of 'd' to be 'e(%d)', but got '%v'", e.value, d.parent)
		}
		if d.color != nodes.Black {
			t.Errorf("expected 'd' color to be 'black' but got '%s'", d.color)
		}
		if d.left != h {
			t.Errorf("expected 'd' left child to be 'h(%d)' but got '%v'", h.value, d.left.value)
		}
		if d.right != nil {
			t.Errorf("expected 'd' right child to be 'nil' but got '%v'", d.right)
		}

		if a.parent != e {
			t.Errorf("expected 'a' parent to be 'e(%d)', but got '%v'", e.value, a.parent)
		}
		if a.color != nodes.Black {
			t.Errorf("expected 'a' color to be 'black' but got '%s'", a.color)
		}
		if a.left != nil {
			t.Errorf("expected 'a' left child to be 'nil' but got '%v'", a.left)
		}
		if a.right != f {
			t.Errorf("expected 'a' right child to be 'f(%d)' but got '%v'", f.value, a.right.value)
		}

		if h.parent != d {
			t.Errorf("expected 'h' parent to be 'd(%d)', but got '%v'", d.value, h.parent)
		}
		if h.color != nodes.Red {
			t.Errorf("expected 'h' color to be 'red' but got '%s'", h.color)
		}
		if h.left != nil {
			t.Errorf("expected 'h' left child to be 'nil' but got %v", h.left)
		}
		if h.right != nil {
			t.Errorf("expected 'h' right child to be 'nil' but got %v", h.right)
		}

		if f.parent != a {
			t.Errorf("expected 'f' parent to be 'a(%d)', but got '%v'", a.value, f.parent)
		}
		if f.color != nodes.Red {
			t.Errorf("expected 'f' color to be 'red' but got %s", f.color)
		}
		if f.left != nil {
			t.Errorf("expected 'f' left child to be nil but got '%v'", f.left)
		}
		if f.right != nil {
			t.Errorf("expected 'f' right child to be nil but got '%v'", f.right)
		}
	})

	// Case 7: Delete on empty tree
	t.Run("Case 7: Delete on empty tree", func(t *testing.T) {
		root = DeleteNode[int](nil, nil)
		if root != nil {
			t.Errorf("expected root to be nil but got '%v'", root)
		}
	})

	// Case 8: Delete on tree with one node
	t.Run("Case 8: Delete on tree with single node", func(t *testing.T) {
		single := Black(10)
		single.left = nil
		single.right = nil
		single = DeleteNode(single, single)
		if single != nil {
			t.Errorf("expected single to be nil but got '%v'", single)
		}
	})
}

func TestInsertFixup(t *testing.T) {
	t.Run("Basic Recoloring", func(t *testing.T) {
		// Create a tree:
		//         10(B)
		//        /    \
		//     5(R)     15(R)
		root := Black(10)
		root.left = Red(5, root)
		root.right = Red(15, root)

		// Insert a new node (requires recoloring)
		n := Red(7, root.left)
		root.left.right = n

		// Perform fixup
		root = insertFixup(root, n)

		// Verify the structure
		if root.color != nodes.Black {
			t.Errorf("expected root color to be Black, but got %s", root.color)
		}
		if root.left.color != nodes.Black || root.right.color != nodes.Black {
			t.Errorf("expected left and right children of root to be Black")
		}
	})

	t.Run("Left Rotation", func(t *testing.T) {
		// Create a tree:
		//         10(B)
		//        /
		//      5(R)
		//       \
		//       7(R)
		root := Black(10)
		root.color = nodes.Black
		root.left = Red(5, root)
		n := Red(7, root.left)
		root.left.right = n

		// Perform fixup
		root = insertFixup(root, n)

		fmx.Printfln("%v", root)

		// Verify the structure
		if root.value != 7 {
			t.Errorf("expected root value to be 7, but got %v", root.value)
		}
		if root.right.value != 10 {
			t.Errorf("expected right child of root to be 10, but got %v", root.right.value)
		}
		if root.left.value != 5 {
			t.Errorf("expected left child of root to be 5, but got %v", root.left.value)
		}
		if root.color != nodes.Black {
			t.Errorf("expected root color to be Black, but got %s", root.color)
		}
		if root.left.color != nodes.Red {
			t.Errorf("expected left child of root to be Red, but got %s", root.left.color)
		}
		if root.right.color != nodes.Red {
			t.Errorf("expected right child of root to be Red, but got %s", root.right.color)
		}
	})

	t.Run("Right Rotation", func(t *testing.T) {
		// Create a tree:
		//         10(B)
		//              \
		//              15(R)
		//             /
		//           12(R)
		root := Root(10)
		root.right = Red(15, root)
		n := Red(12, root.right)
		root.right.left = n

		// Perform fixup
		root = insertFixup(root, n)

		// Verify the structure
		if root.value != 12 {
			t.Errorf("expected root value to be 12, but got %v", root.value)
		}
		if root.right.value != 15 {
			t.Errorf("expected right child of root to be 15, but got %v", root.right.value)
		}
		if root.left.value != 10 {
			t.Errorf("expected left child of root to be 10, but got %v", root.left.value)
		}
		if root.color != nodes.Black {
			t.Errorf("expected root color to be Black, but got %s", root.color)
		}
		if root.left.color != nodes.Red {
			t.Errorf("expected left child of root to be Red, but got %s", root.left.color)
		}
		if root.right.color != nodes.Red {
			t.Errorf("expected right child of root to be Red, but got %s", root.right.color)
		}
	})

	t.Run("Combined Rotations", func(t *testing.T) {
		// Create a tree:
		//         10(B)
		//        /
		//      5(R)
		//       \
		//       7(R)
		root := Root(10)
		root.left = Red(5, root)
		n := Red(7, root.left)
		root.left.right = n
		n.parent = root.left

		// Perform fixup
		root = insertFixup(root, n)

		// Verify the structure
		if root.value != 7 {
			t.Errorf("expected root value to be 7, but got %v", root.value)
		}
		if root.right.value != 10 {
			t.Errorf("expected right child of root to be 10, but got %v", root.right.value)
		}
		if root.left.value != 5 {
			t.Errorf("expected left child of root to be 5, but got %v", root.left.value)
		}
		if root.color != nodes.Black {
			t.Errorf("expected root color to be Black, but got %s", root.color)
		}
		if root.left.color != nodes.Red {
			t.Errorf("expected left child of root to be Red, but got %s", root.left.color)
		}
		if root.right.color != nodes.Red {
			t.Errorf("expected right child of root to be Red, but got %s", root.right.color)
		}
	})

	t.Run("Single-Node Tree", func(t *testing.T) {
		// Create a single-node tree
		root := Root(10)

		// Perform fixup on the root (no changes expected)
		root = insertFixup(root, root)

		// Verify the structure
		if root.value != 10 || root.color != nodes.Black {
			t.Errorf("Unexpected tree structure after fixup on single-node tree")
		}
	})
}

// TestInsertNode tests the InsertNode function.
//
// after an insertion the test should assert
//   - The tree respects the basic binary-tree properties
//   - The tree remains balanced according to the red-black properties
func TestInsertNode(t *testing.T) {
	// Test 1. Empty tree
	var root *Node[int]
	var ok bool
	cmp := comparators.Ordered[int]{}
	var tree = make(map[string]*Node[int])
	t.Run("test node on empty tree", func(t *testing.T) {
		if len(tree) != 0 {
			tree = make(map[string]*Node[int])
		}
		a := NewNode(10, nodes.Black)
		root, ok = InsertNode(root, a, cmp.Compare)
		if !ok {
			t.Errorf("InsertNode returned false while inserting\n%v", a)
		}

		// Verify the structure
		assertValidRoot(t, "a", a, root)
		if a.left != nil {
			t.Errorf("expected left child of node to be nil, but got\n%v", a.left)
		}
		if a.right != nil {
			t.Errorf("expected right child of node to be nil, but got\n%v", a.right)
		}
		tree["a"] = a
	})

	// Test 2. Tree with one node
	t.Run("test insert higher value than root", func(t *testing.T) {
		// ensure last test state on single run
		if len(tree) == 0 {
			tree = make(map[string]*Node[int])
			a := NewNode(10, nodes.Black)
			root, ok = InsertNode(root, a, cmp.Compare)
			if !ok {
				t.Errorf("InsertNode returned false while inserting\n%v", a)
			}
			tree["a"] = a
		}
		b := NewNode(15, nodes.Black)
		root, ok = InsertNode(root, b, cmp.Compare)
		if !ok {
			t.Errorf("InsertNode returned false while inserting\n%v", b)
		}

		// expected structure
		//    a(10)[Root]
		//		     \
		//		   b(15)[R]
		assertValidRoot(t, "a", tree["a"], root)
		if tree["a"].right != b {
			t.Errorf("expected a.right to be 'b(15)', but got\n%v", tree["a"].right)
		}
		if b.parent != tree["a"] {
			t.Errorf("expected b.parent to be 'a(10)', but got\n%v", b.parent)
		}
		if b.color != nodes.Red {
			t.Errorf("expected b.color to be 'Red', but got '%v'", b.color)
		}
		tree["b"] = b
	})

	// Test 3. Tree with two nodes
	t.Run("insert higher value to force unbalance", func(t *testing.T) {
		// ensure last test state on single run
		if len(tree) < 2 {
			tree = make(map[string]*Node[int])
			a := NewNode(10, nodes.Black)
			root, ok = InsertNode(root, a, cmp.Compare)
			if !ok {
				t.Errorf("InsertNode returned false while inserting\n%v", a)
			}
			tree["a"] = a
			b := NewNode(15, nodes.Black)
			root, ok = InsertNode(root, b, cmp.Compare)
			if !ok {
				t.Errorf("InsertNode returned false while inserting\n%v", b)
			}
			tree["b"] = b
		}

		c := NewNode(20, nodes.Black)
		root, ok = InsertNode(root, c, cmp.Compare)
		if !ok {
			t.Errorf("InsertNode returned false while inserting\n%v", c)
		}

		// expected structure
		//        b(15)[Root]
		//        /        \
		// a(10)[R]        c(20)[R]
		a := tree["a"]
		b := tree["b"]
		assertValidRoot(t, "b", b, root)
		if b.left != a {
			t.Errorf("expected b.left to be 'a(10)', but got\n%v", b.left)
		}
		if b.right != c {
			t.Errorf("expected b.right to be 'c(20)', but got\n%v", b.right)
		}

		if a.parent != b {
			t.Errorf("expected a.parent to be 'b(15)', but got\n%v", a.parent)
		}
		if a.color != nodes.Red {
			t.Errorf("expected a.color to be 'Red', but got '%v'", a.color)
		}
		if a.left != nil {
			t.Errorf("expected a.left to be nil, but got\n%v", a.left)
		}
		if a.right != nil {
			t.Errorf("expected a.right to be nil, but got\n%v", a.right)
		}

		if c.parent != b {
			t.Errorf("expected c.parent to be 'b(15)', but got\n%v", c.parent)
		}
		if c.color != nodes.Red {
			t.Errorf("expected c.color to be 'Red', but got '%v'", c.color)
		}
		if c.left != nil {
			t.Errorf("expected c.left to be nil, but got\n%v", c.left)
		}
		if c.right != nil {
			t.Errorf("expected c.right to be nil, but got\n%v", c.right)
		}
		tree["c"] = c
	})

	t.Run("Insert between root and smaller value", func(t *testing.T) {
		// ensure last test state on single run
		if len(tree) < 3 {
			tree = make(map[string]*Node[int])
			a := NewNode(10, nodes.Black)
			root, _ = InsertNode(root, a, cmp.Compare)
			tree["a"] = a
			b := NewNode(15, nodes.Black)
			root, _ = InsertNode(root, b, cmp.Compare)
			tree["b"] = b
			c := NewNode(20, nodes.Black)
			root, _ = InsertNode(root, c, cmp.Compare)
			tree["c"] = c
		}

		d := NewNode(12, nodes.Black)
		root, ok = InsertNode(root, d, cmp.Compare)
		if !ok {
			t.Errorf("InsertNode returned false while inserting\n%v", d)
		}

		// expected structure
		//          b(15)[Root]
		//          /        \
		//   a(10)[B]        c(20)[B]
		//         \
		//         d(12)[R]
		a := tree["a"]
		b := tree["b"]
		c := tree["c"]
		assertValidRoot(t, "b", b, root)
		if b.left != a {
			t.Errorf("expected b.left to be 'a(10)', but got\n%v", b.left)
		}
		if b.right != c {
			t.Errorf("expected b.right to be 'c(20)', but got\n%v", b.right)
		}

		if a.parent != b {
			t.Errorf("expected a.parent to be 'b(15)', but got\n%v", a.parent)
		}
		if a.color != nodes.Black {
			t.Errorf("expected a.color to be 'Black', but got '%v'", a.color)
		}
		if a.left != nil {
			t.Errorf("expected a.left to be nil, but got\n%v", a.left)
		}
		if a.right != d {
			t.Errorf("expected a.right to be d(12), but got\n%v", a.right)
		}

		if c.parent != b {
			t.Errorf("expected c.parent to be 'b(15)', but got\n%v", c.parent)
		}
		if c.color != nodes.Black {
			t.Errorf("expected c.color to be 'Black', but got '%v'", c.color)
		}
		if c.left != nil {
			t.Errorf("expected c.left to be nil, but got\n%v", c.left)
		}
		if c.right != nil {
			t.Errorf("expected c.right to be nil, but got\n%v", c.right)
		}

		if d.parent != a {
			t.Errorf("expected d.parent to be 'a(10)', but got\n%v", d.parent)
		}
		if d.color != nodes.Red {
			t.Errorf("expected d.color to be 'Red', but got '%v'", d.color)
		}
		if d.left != nil {
			t.Errorf("expected d.left to be nil, but got\n%v", d.left)
		}
		if d.right != nil {
			t.Errorf("expected d.right to be nil, but got\n%v", d.right)
		}
		tree["d"] = d
	})

	t.Run("Insert between d and root to force rotations", func(t *testing.T) {
		// ensure last test state on single run
		if len(tree) < 4 {
			tree = make(map[string]*Node[int])
			a := NewNode(10, nodes.Black)
			root, _ = InsertNode(root, a, cmp.Compare)
			tree["a"] = a
			b := NewNode(15, nodes.Black)
			root, _ = InsertNode(root, b, cmp.Compare)
			tree["b"] = b
			c := NewNode(20, nodes.Black)
			root, _ = InsertNode(root, c, cmp.Compare)
			tree["c"] = c
			d := NewNode(12, nodes.Black)
			root, _ = InsertNode(root, d, cmp.Compare)
			tree["d"] = d
		}

		e := NewNode(13, nodes.Black)
		root, ok = InsertNode(root, e, cmp.Compare)
		if !ok {
			t.Errorf("InsertNode returned false while inserting\n%v", e)
		}

		// expected structure
		//                    b(15)[Black]
		//                   /           \
		//  	   d(12)[Black]          c(20)[Black]
		//           /     \
		// a(10)[Red]      e(13)[Red]
		a := tree["a"]
		b := tree["b"]
		c := tree["c"]
		d := tree["d"]
		assertValidRoot(t, "b", b, root)
		if b.left != d {
			t.Errorf("expected b.left to be 'd(12)', but got\n%v", b.left)
		}
		if b.right != c {
			t.Errorf("expected b.right to be 'c(20)', but got\n%v", b.right)
		}

		if d.parent != b {
			t.Errorf("expected d.parent to be 'b(15)', but got\n%v", d.parent)
		}
		if d.color != nodes.Black {
			t.Errorf("expected d.color to be 'Black', but got '%v'", d.color)
		}
		if d.left != a {
			t.Errorf("expected d.left to be 'a(10)', but got\n%v", d.left)
		}
		if d.right != e {
			t.Errorf("expected d.right to be 'e(13)', but got\n%v", d.right)
		}

		if c.parent != b {
			t.Errorf("expected c.parent to be 'b(15)', but got\n%v", c.parent)
		}
		if c.color != nodes.Black {
			t.Errorf("expected c.color to be 'Black', but got '%v'", c.color)
		}
		if c.left != nil {
			t.Errorf("expected c.left to be nil, but got\n%v", c.left)
		}
		if c.right != nil {
			t.Errorf("expected c.right to be nil, but got\n%v", c.right)
		}

		if a.parent != d {
			t.Errorf("expected a.parent to be d(12) but got\n%v", a.parent)
		}
		if a.color != nodes.Red {
			t.Errorf("expected a.color to be 'Red', but got '%v'", a.color)
		}
		if a.left != nil {
			t.Errorf("expected a.left to be nil, but got\n%v", a.left)
		}
		if a.right != nil {
			t.Errorf("expected a.right to be nil, but got\n%v", a.right)
		}

		if e.parent != d {
			t.Errorf("expected e.parent to be d(12) but got\n%v", e.parent)
		}
		if e.color != nodes.Red {
			t.Errorf("expected e.color to be 'Red', but got '%v'", e.color)
		}
		if e.left != nil {
			t.Errorf("expected e.left to be nil, but got\n%v", e.left)
		}
		if e.right != nil {
			t.Errorf("expected e.right to be nil, but got\n%v", e.right)
		}
		tree["e"] = e
	})

	t.Run("insert on right side shouldn't fixup", func(t *testing.T) {
		// ensure last test state on single run
		if len(tree) < 5 {
			tree = make(map[string]*Node[int])
			a := NewNode(10, nodes.Black)
			root, _ = InsertNode(root, a, cmp.Compare)
			tree["a"] = a
			b := NewNode(15, nodes.Black)
			root, _ = InsertNode(root, b, cmp.Compare)
			tree["b"] = b
			c := NewNode(20, nodes.Black)
			root, _ = InsertNode(root, c, cmp.Compare)
			tree["c"] = c
			d := NewNode(12, nodes.Black)
			root, _ = InsertNode(root, d, cmp.Compare)
			tree["d"] = d
			e := NewNode(13, nodes.Black)
			root, _ = InsertNode(root, e, cmp.Compare)
			tree["e"] = e
		}

		f := NewNode(25, nodes.Black)
		root, ok = InsertNode(root, f, cmp.Compare)
		if !ok {
			t.Errorf("InsertNode returned false while inserting\n%v", f)
		}
		// expected structure
		//                    b(15)[Black]
		//                   /           \
		//  	   d(12)[Black]          c(20)[Black]
		//           /     \                      \
		// a(10)[Red]      e(13)[Red]            f(25)[Red]
		a := tree["a"]
		b := tree["b"]
		c := tree["c"]
		d := tree["d"]
		e := tree["e"]
		assertValidRoot(t, "b", b, root)
		if b.left != d {
			t.Errorf("expected b.left to be 'd(12)', but got\n%v", b.left)
		}
		if b.right != c {
			t.Errorf("expected b.right to be 'c(20)', but got\n%v", b.right)
		}

		if d.parent != b {
			t.Errorf("expected d.parent to be 'b(15)', but got\n%v", d.parent)
		}
		if d.color != nodes.Black {
			t.Errorf("expected d.color to be 'Black', but got '%v'", d.color)
		}
		if d.left != a {
			t.Errorf("expected d.left to be 'a(10)', but got\n%v", d.left)
		}
		if d.right != e {
			t.Errorf("expected d.right to be 'e(13)', but got\n%v", d.right)
		}

		if c.parent != b {
			t.Errorf("expected c.parent to be 'b(15)', but got\n%v", c.parent)
		}
		if c.color != nodes.Black {
			t.Errorf("expected c.color to be 'Black', but got '%v'", c.color)
		}
		if c.left != nil {
			t.Errorf("expected c.left to be nil, but got\n%v", c.left)
		}
		if c.right != f {
			t.Errorf("expected c.right to be f(25), but got\n%v", c.right)
		}

		if a.parent != d {
			t.Errorf("expected a.parent to be d(12) but got\n%v", a.parent)
		}
		if a.color != nodes.Red {
			t.Errorf("expected a.color to be 'Red', but got '%v'", a.color)
		}
		if a.left != nil {
			t.Errorf("expected a.left to be nil, but got\n%v", a.left)
		}
		if a.right != nil {
			t.Errorf("expected a.right to be nil, but got\n%v", a.right)
		}

		if e.parent != d {
			t.Errorf("expected e.parent to be d(12) but got\n%v", e.parent)
		}
		if e.color != nodes.Red {
			t.Errorf("expected e.color to be 'Red', but got '%v'", e.color)
		}
		if e.left != nil {
			t.Errorf("expected e.left to be nil, but got\n%v", e.left)
		}
		if e.right != nil {
			t.Errorf("expected e.right to be nil, but got\n%v", e.right)
		}

		if f.parent != c {
			t.Errorf("expected f.parent to be c(20) but got\n%v", f.parent)
		}
		if f.color != nodes.Red {
			t.Errorf("expected f.color to be 'Red', but got '%v'", f.color)
		}
		if f.left != nil {
			t.Errorf("expected f.left to be nil, but got\n%v", f.left)
		}
		if f.right != nil {
			t.Errorf("expected f.right to be nil, but got\n%v", f.right)
		}
	})
}

func TestRotations(t *testing.T) {
	t.Run("Basic root rotations left", func(t *testing.T) {
		// Create a small tree:
		//   x(10)[Root]
		//          \
		//        y(20)[R]
		//			  \
		//            z(30)[B]
		x := Root(10)
		y := Red(20, x)
		z := Black(30, y)
		x.right = y
		y.right = z

		// Perform left rotation on x (shouldn't rotate because
		// there is no right children)
		root := rotateRight(x, x)

		// expect same structure
		assertValidRoot(t, "x", x, root)
		if y.parent != x {
			t.Errorf("Expected y.parent to be 'x(10)', but got %v", y.parent)
		}
		if z.parent != y {
			t.Errorf("Expected z.parent to be 'y(20)', but got %v", z.parent)
		}
		if x.right != y {
			t.Errorf("Expected x.left to be 'y(20)', but got %v", x.left)
		}

		// Perform left rotation on x (should make y the new root)
		root = rotateLeft(x, x)
		// expected structure after rotation: color unbalanced
		//  	       y(20)[Root]
		//             /         \
		//       x(10)[B]       z(30)[B]
		assertValidRoot(t, "y", y, root)
		if y.right != z {
			t.Errorf("Expected y.right to be 'z(30)', but got %v", y.right)
		}
		if z.parent != y {
			t.Errorf("Expected z.parent to be 'y(20)', but got %v", z.parent)
		}
		if z.color != nodes.Black {
			t.Errorf("Expected z.color to be 'Black', but got %v", z.color)
		}
		if z.left != nil {
			t.Errorf("Expected z.left to be nil, but got %v", z.left)
		}
		if z.right != nil {
			t.Errorf("Expected z.right to be nil, but got %v", z.right)
		}
		if y.left != x {
			t.Errorf("Expected y.left to be 'x(10)', but got %v", y.left)
		}
		if x.parent != y {
			t.Errorf("Expected x.parent to be 'y(20)', but got %v", x.parent)
		}
		if x.color != nodes.Black {
			t.Errorf("Expected x.color to be 'Black', but got %v", x.color)
		}
		if x.left != nil {
			t.Errorf("Expected x.left to be nil, but got %v", x.left)
		}
		if x.right != nil {
			t.Errorf("Expected x.right to be nil, but got %v", x.right)
		}
	})

	t.Run("Basic root rotations right", func(t *testing.T) {
		// Create a small tree:
		//  	       x(30)[Root]
		//                /
		//        y(20)[R]
		//			/
		//   z(10)[B]
		x := Root(30)
		y := Red(20, x)
		z := Black(10, y)
		x.left = y
		y.left = z

		// Perform left rotation on x (shouldn't rotate because
		// there is no right children)
		root := rotateLeft(x, x)

		// expect same structure
		assertValidRoot(t, "x", x, root)
		if y.parent != x {
			t.Errorf("Expected y.parent to be 'x(10)', but got %v", y.parent)
		}
		if z.parent != y {
			t.Errorf("Expected z.parent to be 'y(20)', but got %v", z.parent)
		}
		if x.left != y {
			t.Errorf("Expected x.left to be 'y(20)', but got %v", x.left)
		}

		// Perform right rotation on x (should make y the new root)
		root = rotateRight(x, x)
		// expected structure after rotation: color unbalanced
		//  	       y(20)[Root]
		//             /         \
		//       z(30)[B]       x(10)[B]
		assertValidRoot(t, "y", y, root)
		if y.left != z {
			t.Errorf("Expected y.left to be 'z(30)', but got %v", y.left)
		}
		if z.parent != y {
			t.Errorf("Expected z.parent to be 'y(20)', but got %v", z.parent)
		}
		if z.color != nodes.Black {
			t.Errorf("Expected z.color to be 'Black', but got %v", z.color)
		}
		if z.left != nil {
			t.Errorf("Expected z.left to be nil, but got %v", z.left)
		}
		if z.right != nil {
			t.Errorf("Expected z.right to be nil, but got %v", z.right)
		}
		if y.right != x {
			t.Errorf("Expected y.right to be 'x(10)', but got %v", y.right)
		}
		if x.parent != y {
			t.Errorf("Expected x.parent to be 'y(20)', but got %v", x.parent)
		}
		if x.color != nodes.Black {
			t.Errorf("Expected x.color to be 'Black', but got %v", x.color)
		}
		if x.left != nil {
			t.Errorf("Expected x.left to be nil, but got %v", x.left)
		}
		if x.right != nil {
			t.Errorf("Expected x.right to be nil, but got %v", x.right)
		}
	})

	t.Run("Nil root", func(t *testing.T) {
		// Create a node with nil root
		x := Black(10)

		// Perform left rotation on x (no right child)
		root := rotateLeft(nil, x)

		// Verify that the tree remains unchanged
		if root != nil {
			t.Errorf("Expected root to be nil, but got %v", root)
		}
		if x.left != nil || x.right != nil {
			t.Errorf("Expected x to have no children, but got left=%v, right=%v", x.left, x.right)
		}

		root = rotateRight(nil, x)

		// Verify that the tree remains unchanged
		if root != nil {
			t.Errorf("Expected root to be nil, but got %v", root)
		}
		if x.left != nil || x.right != nil {
			t.Errorf("Expected x to have no children, but got \nleft: '%v'\nright: %v",
				x.left, x.right)
		}
	})

	t.Run("left rotation: large balanced tree", func(t *testing.T) {
		// Create a larger tree:
		//         a(10)[Root]
		//           /    \
		//     b(5)[B]   c(20)[R]
		//                /    \
		//         d(15)[B]    e(30)[B]
		a := Root(10)
		a.left = Black(5, a)
		b := a.left
		a.right = Red(20, a)
		c := a.right
		c.left = Black(15, c)
		d := c.left
		c.right = Black(30, c)
		e := c.right
		root := a

		// Test 1: rotate on root
		root = rotateLeft(root, a)
		// expected structure (color unbalance):
		//         c(20)[Root]
		//           /    \
		//     a(10)[B]   e(30)[B]
		//      /     \
		// b(5)[B]    d(15)[B]
		assertValidRoot(t, "c", c, root)
		if c.left != a {
			t.Errorf("Expected c.left to be a(10), but got:\n%v", c.left)
		}
		if c.right != e {
			t.Errorf("Expected c.right to be e(30), but got\n%v", c.right)
		}

		if a.parent != c {
			t.Errorf("Expected a.parent to be c(20), but got\n%v", a.parent)
		}
		if a.color != nodes.Black {
			t.Errorf("Expected a.color to be 'Black', but got '%v'", a.color)
		}
		if a.left != b {
			t.Errorf("Expected a.left to be b(5), but got\n%v", a.left)
		}
		if a.right != d {
			t.Errorf("Expected a.right to be d(15), but got\n%v", a.right)
		}

		if e.parent != c {
			t.Errorf("Expected e.parent to be c(20), but got\n%v", e.parent)
		}
		if e.color != nodes.Black {
			t.Errorf("Expected e.color to be 'Black', but got '%v'", e.color)
		}
		if e.left != nil {
			t.Errorf("Expected e.left to be nil, but got\n%v", e.left)
		}
		if e.right != nil {
			t.Errorf("Expected e.right to be nil, but got\n%v", e.right)
		}

		if b.parent != a {
			t.Errorf("Expected b.parent to be a(10), but got\n%v", b.parent)
		}
		if b.color != nodes.Black {
			t.Errorf("Expected b.color to be 'Black', but got '%v'", b.color)
		}
		if b.left != nil {
			t.Errorf("Expected b.left to be nil, but got\n%v", b.left)
		}
		if b.right != nil {
			t.Errorf("Expected b.right to be nil, but got\n%v", b.right)
		}

		if d.parent != a {
			t.Errorf("Expected d.parent to be a(10), but got\n%v", d.parent)
		}
		if d.color != nodes.Black {
			t.Errorf("Expected d.color to be 'Black', but got '%v'", d.color)
		}
		if d.left != nil {
			t.Errorf("Expected d.left to be nil, but got\n%v", d.left)
		}
		if d.right != nil {
			t.Errorf("Expected d.right to be nil, but got\n%v", d.right)
		}

		// Test 2: Rotate on root to left
		root = rotateRight(root, c)
		// Expect the tree becomes how it was (but with c black):
		//         a(10)[Root]
		//           /    \
		//     b(5)[B]   c(20)[B]
		//                /    \
		//         d(15)[B]    e(30)[B]
		assertValidRoot(t, "a", a, root)
		if a.left != b {
			t.Errorf("Expected a.left to be b(5), but got\n%v", a.left)
		}
		if a.right != c {
			t.Errorf("Expected a.right to be c(20), but got\n%v", a.right)
		}

		if b.parent != a {
			t.Errorf("Expected b.parent to be a(10), but got\n%v", b.parent)
		}
		if b.color != nodes.Black {
			t.Errorf("Expected b.color to be 'Black', but got '%v'", b.color)
		}
		if b.left != nil {
			t.Errorf("Expected b.left to be nil, but got\n%v", b.left)
		}
		if b.right != nil {
			t.Errorf("Expected b.right to be nil, but got\n%v", b.right)
		}

		if c.parent != a {
			t.Errorf("Expected c.parent to be a(10), but got\n%v", c.parent)
		}
		if c.color != nodes.Black {
			t.Errorf("Expected c.color to be 'Black', but got '%v'", c.color)
		}
		if c.left != d {
			t.Errorf("Expected c.left to be d(15), but got\n%v", c.left)
		}
		if c.right != e {
			t.Errorf("Expected c.right to be e(30), but got\n%v", c.right)
		}

		if d.parent != c {
			t.Errorf("Expected d.parent to be c(20), but got\n%v", d.parent)
		}
		if d.color != nodes.Black {
			t.Errorf("Expected d.color to be 'Black', but got '%v'", d.color)
		}
		if d.left != nil {
			t.Errorf("Expected d.left to be nil, but got\n%v", d.left)
		}
		if d.right != nil {
			t.Errorf("Expected d.right to be nil, but got\n%v", d.right)
		}

		if e.parent != c {
			t.Errorf("Expected e.parent to be c(20), but got\n%v", e.parent)
		}
		if e.color != nodes.Black {
			t.Errorf("Expected e.color to be 'Black', but got '%v'", e.color)
		}
		if e.left != nil {
			t.Errorf("Expected e.left to be nil, but got\n%v", e.left)
		}
		if e.right != nil {
			t.Errorf("Expected e.right to be nil, but got\n%v", e.right)
		}

		// Test 3: Rotate root to right
		root = rotateRight(root, a)
		// Expect partial flattening to:
		//   b(5)[Root]
		//        \
		//     a(10)[B]
		//          \
		//        c(20)[B]
		//	      /     \
		//  d(15)[B]    e(30)[B]
		assertValidRoot(t, "b", b, root)
		if b.left != nil {
			t.Errorf("Expected b.left to be nil, but got\n%v", b.left)
		}
		if b.right != a {
			t.Errorf("Expected b.right to be a(10), but got\n%v", b.right)
		}

		if a.parent != b {
			t.Errorf("Expected a.parent to be b(5), but got\n%v", a.parent)
		}
		if a.color != nodes.Black {
			t.Errorf("Expected a.color to be 'Black', but got '%v'", a.color)
		}
		if a.left != nil {
			t.Errorf("Expected a.left to be nil, but got\n%v", a.left)
		}
		if a.right != c {
			t.Errorf("Expected a.right to be c(20), but got\n%v", a.right)
		}

		if c.parent != a {
			t.Errorf("Expected c.parent to be a(10), but got\n%v", c.parent)
		}
		if c.color != nodes.Black {
			t.Errorf("Expected c.color to be 'Black', but got '%v'", c.color)
		}
		if c.left != d {
			t.Errorf("Expected c.left to be d(15), but got\n%v", c.left)
		}
		if c.right != e {
			t.Errorf("Expected c.right to be e(30), but got\n%v", c.right)
		}

		if d.parent != c {
			t.Errorf("Expected d.parent to be c(20), but got\n%v", d.parent)
		}
		if d.color != nodes.Black {
			t.Errorf("Expected d.color to be 'Black', but got '%v'", d.color)
		}
		if d.left != nil {
			t.Errorf("Expected d.left to be nil, but got\n%v", d.left)
		}
		if d.right != nil {
			t.Errorf("Expected d.right to be nil, but got\n%v", d.right)
		}

		if e.parent != c {
			t.Errorf("Expected e.parent to be c(20), but got\n%v", e.parent)
		}
		if e.color != nodes.Black {
			t.Errorf("Expected e.color to be 'Black', but got '%v'", e.color)
		}
		if e.left != nil {
			t.Errorf("Expected e.left to be nil, but got\n%v", e.left)
		}
		if e.right != nil {
			t.Errorf("Expected e.right to be nil, but got\n%v", e.right)
		}

		// test 4: rotate on root.right.right (c(20))
		root = rotateRight(root, c)
		// Expect full left flattening to:
		//   b(5)[Root]
		//        \
		//     a(10)[B]
		//          \
		//        d(15)[B]
		//            \
		//          c(20)[B]
		//	            \
		//            e(30)[B]
		assertValidRoot(t, "b", b, root)
		if b.left != nil {
			t.Errorf("Expected b.left to be nil, but got\n%v", b.left)
		}
		if b.right != a {
			t.Errorf("Expected b.right to be a(10), but got\n%v", b.right)
		}

		if a.parent != b {
			t.Errorf("Expected a.parent to be b(5), but got\n%v", a.parent)
		}
		if a.color != nodes.Black {
			t.Errorf("Expected a.color to be 'Black', but got '%v'", a.color)
		}
		if a.left != nil {
			t.Errorf("Expected a.left to be nil, but got\n%v", a.left)
		}
		if a.right != d {
			t.Errorf("Expected a.right to be d(15), but got\n%v", a.right)
		}

		if d.parent != a {
			t.Errorf("Expected d.parent to be a(10), but got\n%v", d.parent)
		}
		if d.color != nodes.Black {
			t.Errorf("Expected d.color to be 'Black', but got '%v'", d.color)
		}
		if d.left != nil {
			t.Errorf("Expected d.left to be c(20), but got\n%v", d.left)
		}
		if d.right != c {
			t.Errorf("Expected d.right to be e(30), but got\n%v", d.right)
		}

		if c.parent != d {
			t.Errorf("Expected c.parent to be d(15), but got\n%v", c.parent)
		}
		if c.color != nodes.Black {
			t.Errorf("Expected c.color to be 'Black', but got '%v'", c.color)
		}
		if c.left != nil {
			t.Errorf("Expected c.left to be nil, but got\n%v", c.left)
		}
		if c.right != e {
			t.Errorf("Expected c.right to be nil, but got\n%v", c.right)
		}

		if e.parent != c {
			t.Errorf("Expected e.parent to be c(20), but got\n%v", e.parent)
		}
		if e.color != nodes.Black {
			t.Errorf("Expected e.color to be 'Black', but got '%v'", e.color)
		}
		if e.left != nil {
			t.Errorf("Expected e.left to be nil, but got\n%v", e.left)
		}
		if e.right != nil {
			t.Errorf("Expected e.right to be nil, but got\n%v", e.right)
		}

		// Test 5: rotate on root.right.right.right.right to right (e)
		root = rotateRight(root, e)
		// Expect no change:
		assertValidRoot(t, "b", b, root)
		if b.left != nil {
			t.Errorf("Expected b.left to be nil, but got\n%v", b.left)
		}
		if b.right != a {
			t.Errorf("Expected b.right to be a(10), but got\n%v", b.right)
		}

		if a.parent != b {
			t.Errorf("Expected a.parent to be b(5), but got\n%v", a.parent)
		}
		if a.color != nodes.Black {
			t.Errorf("Expected a.color to be 'Black', but got '%v'", a.color)
		}
		if a.left != nil {
			t.Errorf("Expected a.left to be nil, but got\n%v", a.left)
		}
		if a.right != d {
			t.Errorf("Expected a.right to be d(15), but got\n%v", a.right)
		}

		if d.parent != a {
			t.Errorf("Expected d.parent to be a(10), but got\n%v", d.parent)
		}
		if d.color != nodes.Black {
			t.Errorf("Expected d.color to be 'Black', but got '%v'", d.color)
		}
		if d.left != nil {
			t.Errorf("Expected d.left to be c(20), but got\n%v", d.left)
		}
		if d.right != c {
			t.Errorf("Expected d.right to be e(30), but got\n%v", d.right)
		}

		if c.parent != d {
			t.Errorf("Expected c.parent to be d(15), but got\n%v", c.parent)
		}
		if c.color != nodes.Black {
			t.Errorf("Expected c.color to be 'Black', but got '%v'", c.color)
		}
		if c.left != nil {
			t.Errorf("Expected c.left to be nil, but got\n%v", c.left)
		}
		if c.right != e {
			t.Errorf("expected c.right to be e(30), but got\n%v", c.right)
		}

		if e.parent != c {
			t.Errorf("Expected e.parent to be c(20), but got\n%v", e.parent)
		}
		if e.color != nodes.Black {
			t.Errorf("Expected e.color to be 'Black', but got '%v'", e.color)
		}
		if e.left != nil {
			t.Errorf("Expected e.left to be nil, but got\n%v", e.left)
		}
		if e.right != nil {
			t.Errorf("Expected e.right to be nil, but got\n%v", e.right)
		}

		// Test 6: rotate on root.right.right.right to left (c(20))
		root = rotateLeft(root, c)
		// expected structure:
		//   b(5)[Root]
		//        \
		//     a(10)[B]
		//          \
		//        d(15)[B]
		//            \
		//          e(30)[B]
		//           /
		//         c(20)[B]
		assertValidRoot(t, "b", b, root)
		if b.left != nil {
			t.Errorf("Expected b.left to be nil, but got\n%v", b.left)
		}
		if b.right != a {
			t.Errorf("Expected b.right to be a(10), but got\n%v", b.right)
		}

		if a.parent != b {
			t.Errorf("Expected a.parent to be b(5), but got\n%v", a.parent)
		}
		if a.color != nodes.Black {
			t.Errorf("Expected a.color to be 'Black', but got '%v'", a.color)
		}
		if a.left != nil {
			t.Errorf("Expected a.left to be nil, but got\n%v", a.left)
		}
		if a.right != d {
			t.Errorf("Expected a.right to be d(15), but got\n%v", a.right)
		}

		if d.parent != a {
			t.Errorf("Expected d.parent to be a(10), but got\n%v", d.parent)
		}
		if d.color != nodes.Black {
			t.Errorf("Expected d.color to be 'Black', but got '%v'", d.color)
		}
		if d.left != nil {
			t.Errorf("Expected d.left to be nil, but got\n%v", d.left)
		}
		if d.right != e {
			t.Errorf("Expected d.right to be e(30), but got\n%v", d.right)
		}

		if e.parent != d {
			t.Errorf("Expected e.parent to be d(15), but got\n%v", e.parent)
		}
		if e.color != nodes.Black {
			t.Errorf("Expected e.color to be 'Black', but got '%v'", e.color)
		}
		if e.left != c {
			t.Errorf("Expected e.left to be c(20), but got\n%v", e.left)
		}
		if e.right != nil {
			t.Errorf("Expected e.right to be nil, but got\n%v", e.right)
		}

		if c.parent != e {
			t.Errorf("Expected c.parent to be e(30), but got\n%v", c.parent)
		}
		if c.color != nodes.Black {
			t.Errorf("Expected c.color to be 'Black', but got '%v'", c.color)
		}
		if c.left != nil {
			t.Errorf("Expected c.left to be nil, but got\n%v", c.left)
		}
		if c.right != nil {
			t.Errorf("Expected c.right to be nil, but got\n%v", c.right)
		}
	})
}
