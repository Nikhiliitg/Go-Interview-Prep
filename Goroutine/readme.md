** Go  – Interview Q&A **

# Go WaitGroup Interview Guide – Q&A (Including Trick Questions)

> Complete guide to `sync.WaitGroup` in Go. All answers are concise and interview-ready.

---

## 1. What is a WaitGroup?

**Q:** What is a WaitGroup in Go?  
**A:** A `WaitGroup` is a synchronization primitive from Go's `sync` package. It allows the main goroutine to wait until a set of goroutines finish. Analogy: A teacher waits until all kids finish homework before leaving.

---

## 2. Why do we use WaitGroup?

**Q:** Why is WaitGroup necessary?  
**A:** Goroutines run independently and do not block `main()`. Without a WaitGroup, `main()` may exit before workers finish. WaitGroup ensures proper synchronization.

---

## 3. Common Mistakes & Effects

**Q:** What are common mistakes with WaitGroup?  

**A:**  
1. **Forgetting Done()** → `Wait()` blocks forever (deadlock).  
2. **Calling Wait() inside a goroutine** → can cause deadlocks or race conditions.  
3. **Calling Add() after launching goroutines** → panic or unexpected behavior.  
4. **Calling Done() more times than Add()** → panic.  
5. **Reusing WaitGroup before counter reaches zero** → race condition or undefined behavior.

**Notes:** Always call `Add(n)` before launching goroutines, `Done()` exactly once inside each goroutine, and `Wait()` only in the goroutine that needs to wait (usually `main()`).

---

## 4. Correct Idiomatic Pattern

**Q:** What is the correct way to use WaitGroup?  
**A:**  
- Call `Add(n)` **before** launching goroutines.  
- Inside each goroutine, call `defer Done()` to guarantee completion.  
- Call `Wait()` in the main goroutine to block until all goroutines finish.

---

## 5. Elevator Pitch Answer

**Q:** How would you explain WaitGroup in 15 seconds?  

**A:**  
“A WaitGroup tracks goroutines with a counter. `Add()` increments, `Done()` decrements, `Wait()` blocks until zero. Idiomatic pattern: Add before launching goroutines, defer Done inside each, Wait in main. Common pitfalls: forgetting Done (deadlock), Add after launch (panic/race), calling Wait inside a goroutine (deadlock).”

---

## 6. Trick Questions Interviewers May Ask

**⚡ Trick Q1: What happens if you forget to call Done() in a goroutine?**  
**Answer:**  
If a goroutine never calls `Done()`, the counter never decreases. `Wait()` will block forever (deadlock).  
**Pro tip:** “Use `defer wg.Done()` at the start of the goroutine to guarantee it runs.”

**⚡ Trick Q2: What happens if you call Wait() inside a goroutine?**  
**Answer:**  
`Wait()` should almost always be called in the main goroutine. Calling it inside another goroutine while adding tasks can cause race or deadlock.  
**Pro tip:** “Call `Add()` before launching goroutines, defer `Done()` inside them, and call `Wait()` only once, in the goroutine that needs to wait.”

**⚡ Trick Q3: What if you call Add() after launching goroutines?**  
**Answer:**  
Unsafe — the goroutine could finish and call `Done()` before `Add()` is registered, causing a panic.  
**Pro tip:** “Always call `Add(n)` before starting goroutines.”

**⚡ Trick Q4: Can we reuse a WaitGroup?**  
**Answer:**  
Yes, but only after the counter has reached zero. It can then be reused for another batch of goroutines.

**⚡ Trick Q5: What if Done() is called more times than Add()?**  
**Answer:**  
It causes a panic: `sync: negative WaitGroup counter`.

**Interview Tip Summary:**  
“A WaitGroup tracks goroutines with a counter. `Add()` increases it, `Done()` decreases it, and `Wait()` blocks until zero. Common pitfalls: forgetting `Done()` (deadlock) or misplacing `Add()` (race condition). Idiomatic pattern: `Add()` before goroutines, `defer Done()` inside, `Wait()` once in main.”

---

## 7. Pro Tips for Interview

1. Always `Add(n)` **before** starting goroutines.  
2. Use `defer wg.Done()` inside each goroutine.  
3. Call `Wait()` **only once**, usually in `main()`.  
4. Reuse WaitGroup only after the counter reaches zero.  
5. Think like a teacher: wait for all kids (goroutines) before leaving the classroom.

---
