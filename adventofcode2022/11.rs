
fn star1() {
    use std::collections::VecDeque;

    struct Monkey {
        items : VecDeque<i32>,
        optype : char,
        opother : Option<i32>,
        divisible_by : i32,
        throw_to_true : i32,
        throw_to_false : i32,
        inspects : i32,
    }

    impl Monkey {
        fn new(s : &str) -> Monkey {
            let mut items = VecDeque::new();
            let lines = s.split("\n").collect::<Vec<&str>>();
            for item in lines[1]["  Starting items: ".len()..].split(", ") {
                items.push_back(item.parse::<i32>().unwrap());
            }
            let opeq = &lines[2]["  Operation: new = old ".len()..];
            let optype = opeq.chars().nth(0).unwrap();
            let opother = opeq[2..].parse().ok();
            let divisible_by = lines[3]["  Test: divisible by ".len()..].parse().unwrap();
            let throw_to_true = lines[4]["    If true: throw to monkey ".len()..].parse().unwrap();
            let throw_to_false = lines[5]["    If false: throw to monkey ".len()..].parse().unwrap();
            return Monkey { items, optype, opother, divisible_by, throw_to_true, throw_to_false, inspects:0 };
        }

        fn op(&self, wl : i32) -> i32 {
            let other = self.opother.unwrap_or(wl);
            if self.optype == '+' { return wl + other} else { return wl * other; }
        }

        fn next(&mut self) -> Option<i32> {
            let out = self.items.pop_front();
            if let Some(_) = out { self.inspects += 1; }
            return out;
        }

        fn throw_to(&self, worry:i32) -> usize {
            if worry % self.divisible_by == 0 { return self.throw_to_true as usize; } { return self.throw_to_false as usize; }
        }

        fn add(&mut self, worry : i32) {
            self.items.push_back(worry);
        }
    }

    let mut monkeys = Vec::new();
    
    for lines in std::fs::read_to_string("11.txt").unwrap().split("\n\n") {
        //println!("## {lines}");
        monkeys.push(Monkey::new(lines));
    }

    for _ in 1..21 {
        for turn in 0..monkeys.len() {
            //println!("Monkey {turn}:");
            loop {
                let monkey = &mut monkeys[turn];
                let next = monkey.next();
                if next == None { break; }
                let worry1 = next.unwrap();
                let worry2 = monkey.op(worry1);
                let worry3 = worry2 / 3;
                let throw_to = monkey.throw_to(worry3);
                //println!("{worry1} {worry2} {worry3} -> {throw_to}");
                let throw_to_monkey = &mut monkeys[throw_to];
                throw_to_monkey.add(worry3);
            }
        }
    }

    let mut inspects = Vec::new();
    for monkey in monkeys {
        inspects.push(monkey.inspects);
    }
    inspects.sort();
    println!("{}", inspects[inspects.len()-1] * inspects[inspects.len()-2]);
}

fn star2() {
    use std::collections::VecDeque;

    struct Monkey {
        items : VecDeque<i128>,
        optype : char,
        opother : Option<i128>,
        divisible_by : i128,
        throw_to_true : i128,
        throw_to_false : i128,
        inspects : i128,
    }

    impl Monkey {
        fn new(s : &str) -> Monkey {
            let mut items = VecDeque::new();
            let lines = s.split("\n").collect::<Vec<&str>>();
            for item in lines[1]["  Starting items: ".len()..].split(", ") {
                items.push_back(item.parse::<i128>().unwrap());
            }
            let opeq = &lines[2]["  Operation: new = old ".len()..];
            let optype = opeq.chars().nth(0).unwrap();
            let opother = opeq[2..].parse().ok();
            let divisible_by = lines[3]["  Test: divisible by ".len()..].parse().unwrap();
            let throw_to_true = lines[4]["    If true: throw to monkey ".len()..].parse().unwrap();
            let throw_to_false = lines[5]["    If false: throw to monkey ".len()..].parse().unwrap();
            return Monkey { items, optype, opother, divisible_by, throw_to_true, throw_to_false, inspects:0 };
        }

        fn op(&self, wl : i128) -> i128 {
            let other = self.opother.unwrap_or(wl);
            if self.optype == '+' { return wl + other} else { return wl * other; }
        }

        fn next(&mut self) -> Option<i128> {
            let out = self.items.pop_front();
            if let Some(_) = out { self.inspects += 1; }
            return out;
        }

        fn throw_to(&self, worry:i128) -> usize {
            if worry % self.divisible_by == 0 { return self.throw_to_true as usize; } { return self.throw_to_false as usize; }
        }

        fn add(&mut self, worry : i128) {
            self.items.push_back(worry);
        }
    }

    let mut monkeys = Vec::new();
    
    for lines in std::fs::read_to_string("11.txt").unwrap().split("\n\n") {
        //println!("## {lines}");
        monkeys.push(Monkey::new(lines));
    }

    let mut divall = 1;
    for monkey in &monkeys {
        divall *= monkey.divisible_by;
    }
    
    for _round in 1..10001 {
        for turn in 0..monkeys.len() {
            //println!("Monkey {turn}:");
            loop {
                let monkey = &mut monkeys[turn];
                let next = monkey.next();
                if next == None { break; }
                let worry1 = next.unwrap();
                let worry2 = monkey.op(worry1);
                let worry3 = worry2 % divall;// / 3;
                let throw_to = monkey.throw_to(worry3);
                //println!("{worry1} {worry2} {worry3} -> {throw_to}");
                let throw_to_monkey = &mut monkeys[throw_to];
                throw_to_monkey.add(worry3);
            }
        }
        // if _round % 1000 == 0 {
        //     println!("{_round}");
        //     for i in 0..monkeys.len() {
        //         println!("  {i} = {:?}", monkeys[i].inspects);
        //     }
        // }
    }

    let mut inspects = Vec::new();
    for monkey in monkeys {
        inspects.push(monkey.inspects);
    }
    inspects.sort();
    println!("{}", inspects[inspects.len()-1] * inspects[inspects.len()-2]);
    // 14103881622 too low
}

fn main() {
    star1();
    star2();
}
