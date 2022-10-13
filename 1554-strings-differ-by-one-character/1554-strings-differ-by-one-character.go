// 1. set + memoization
// 2. trie + dfs
// 3. rabin-karp <

func differByOne(dict []string) bool {
    return withRabinKarp(dict)
    // return withTrie(dict)
    // return viaSetÅndMemo(dict)
}

var base int = 27
var mod int = 100_007
func withRabinKarp(dict []string) bool {
    hashes := make([]int, len(dict))
    for i, str := range dict {
        hashes[i] = rabinFingerprint(str)
    }
    
    for j,pow:=len(dict[0])-1,1; j>=0; j-- {
        seen := make(map[int][]int)
        for i, w := range dict {
            // i=0 w=abcd
            // h = h(bcd) = h(abcd) - h(a*(base*0))
            // i=1 w=abcd
            // h = h(acd) = h(abcd) - h(b*(base*1))
			h := (mod + hashes[i] - pow * int(w[j]-'a') % mod) % mod
			for _, idx := range seen[h] {
                // i=1 seen=abcd w=aacd
                // a=a & cd=cd
				if w[:j] == dict[idx][:j] && w[j+1:] == dict[idx][j+1:] {
					return true
				}
			}
			seen[h] = append(seen[h], i)
		}
		pow = (pow * base % mod)
	}
	return false
}

func rabinFingerprint(s string) int {
    // f(x) = M0 + M1*X + ... + Mn-1*X^n-1
    // X > man(Mi) -> base
    hash := 0
    for _, ch := range s {
        hash = (hash*base + int(ch-'a')) % mod
    }
    return hash
}

// O(n*m)
// 57%
func withTrie(dict []string) bool {
    trie := NewTrie()
    
    // 문자열을 trie에 넣기 전에 trie를 dfs로 검색하면서
    // 조건에 부합하는 문자열이 이미 있는지 확인해본다.
    for _, str := range dict {
        if dfs(str, *trie.root, 0, 0) {
            return true
        }
        
        tmp := trie.root
        for _, ch := range str {
            idx := ch - 'a'
            if tmp.children[idx] == nil {
                tmp.children[idx] = NewTrieNode(byte(ch))
            }
            tmp = tmp.children[idx]
        }
    }
    
    return false
}

func dfs(str string, node TrieNode, idx, diffCnt int) bool {
    if diffCnt > 1 {
        return false
    }
    
    if idx == len(str) {
        return true
    }
    
    for i := range node.children {
        if node.children[i] != nil {
            if str[idx] == node.children[i].c && dfs(str, *node.children[i], idx+1, diffCnt) {
                return true
            } else if dfs(str, *node.children[i], idx+1, diffCnt+1) {
                return true
            }
        }
    }
    
    return false
}

type TrieNode struct {
    c byte
    children [26]*TrieNode
}

func NewTrieNode(c byte) *TrieNode {
    node := &TrieNode{
        c: c,
    }
    for i:=0; i<26; i++ {
        node.children[i] = nil
    }
    return node
}

type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie {
    root := NewTrieNode('0')
    return &Trie{
        root,
    }
}

// O(n*m)
// out of memory
func viaSetÅndMemo(dict []string) bool {
    set := make(map[string]int)
    
    // dict 내의 문자열들을 순회하면서
    // i번째 위치에 있는 문자를 제외한 문자열이 set에 존재하는지 확인하고
    // 있다면 true, 없다면 현재 문자열의 index와 함께 set에 넣는다.
    // e.g., abcd -> 0bcd a0cd ab0d abc0
    for c, s := range dict {
        for i := range s {
            tmp := s[:i] + "0" + s[i+1:]
            
            if idx, ok := set[tmp]; ok {
                if idx != c {
                    return true
                }
            }
            set[tmp] = c
        }
    }
    
    return false
}
