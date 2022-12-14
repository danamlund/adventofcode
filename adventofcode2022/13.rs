
fn star1() {
    use std::collections::VecDeque;

    #[derive(Debug)]
    #[derive(Copy, Clone)]
    enum Packet { L, R, N(i32), }

    #[derive(Debug)]
    struct Pair { l : VecDeque<Packet>, r : VecDeque<Packet> }

    fn parse_line(line : &str) -> VecDeque<Packet> {
        let mut packets = VecDeque::new();
        let mut cs = String::new();
        for c in line.chars() {
            if c.is_numeric() {
                cs.push(c);
            } else {
                if cs.len() >= 1 {
                    packets.push_back(Packet::N(cs.parse::<i32>().unwrap()));
                    cs.clear();
                }
                if c == '[' { packets.push_back(Packet::L); }
                if c == ']' { packets.push_back(Packet::R); }
            }
        }
        if cs.len() >= 1 {
            packets.push_back(Packet::N(cs.parse::<i32>().unwrap()));
            cs.clear();
        }
        return packets;
    }
    
    let mut pairs : Vec<Pair> = Vec::new();
    for pair_str in std::fs::read_to_string("13.txt").unwrap().split("\n\n") {
        let lines = pair_str.split("\n").collect::<Vec<_>>();
        pairs.push(Pair { l : parse_line(lines[0]), r : parse_line(lines[1]) });
    }

    fn cmp(pair : &mut Pair) -> i32 {
        return cmp1(&mut pair.l, &mut pair.r);
    }
    fn cmp1(l : &mut VecDeque<Packet>, r : &mut VecDeque<Packet>) -> i32 {
        loop {
            let le = l.pop_front().unwrap();
            let re = r.pop_front().unwrap();
            //println!("? {:?}, {:?}", le, re);
            match (le, re) {
                (Packet::N(ln), Packet::L) => {
                    l.push_front(Packet::R);
                    l.push_front(Packet::N(ln));
                    l.push_front(Packet::L);
                    r.push_front(re);
                },
                (Packet::L, Packet::N(rn)) => {
                    r.push_front(Packet::R);
                    r.push_front(Packet::N(rn));
                    r.push_front(Packet::L);
                    l.push_front(le);
                },
                (Packet::N(ln), Packet::N(lr)) => {
                    if ln < lr {
                        return -1;
                    } else if ln > lr {
                        return 1;
                    } else {
                        // continue
                    }
                },
                (Packet::L, Packet::L) => {
                    let c = cmp1(l, r);
                    if c != 0 { return c; }
                },
                (Packet::R, Packet::R) => {
                    return 0;
                },
                (Packet::R, _) => {
                    return -1;
                },
                (_, Packet::R) => {
                    return 1;
                },
            }
        }
    }

    let mut sum = 0;
    let mut i = 0;
    for mut pair in pairs {
        //println!("## {:?}", pair);
        if cmp(&mut pair) == -1 {
            sum += i+1;
        }
        i += 1;
    }
    
    println!("{}", sum);
}

fn star2() {
    use std::collections::VecDeque;
    use std::cmp::Ordering;

    #[derive(Debug)]
    #[derive(Copy, Clone, PartialEq)]
    enum Packet { L, R, N(i32), }

    #[derive(Debug)]
    struct Pair { l : VecDeque<Packet>, r : VecDeque<Packet> }

    fn parse_line(line : &str) -> VecDeque<Packet> {
        let mut packets = VecDeque::new();
        let mut cs = String::new();
        for c in line.chars() {
            if c.is_numeric() {
                cs.push(c);
            } else {
                if cs.len() >= 1 {
                    packets.push_back(Packet::N(cs.parse::<i32>().unwrap()));
                    cs.clear();
                }
                if c == '[' { packets.push_back(Packet::L); }
                if c == ']' { packets.push_back(Packet::R); }
            }
        }
        if cs.len() >= 1 {
            packets.push_back(Packet::N(cs.parse::<i32>().unwrap()));
            cs.clear();
        }
        return packets;
    }
    
    let mut pairs : Vec<Pair> = Vec::new();
    for pair_str in std::fs::read_to_string("13.txt").unwrap().split("\n\n") {
        let lines = pair_str.split("\n").collect::<Vec<_>>();
        pairs.push(Pair { l : parse_line(lines[0]), r : parse_line(lines[1]) });
    }

    fn cmp1(l : &mut VecDeque<Packet>, r : &mut VecDeque<Packet>) -> Ordering {
        loop {
            let le = l.pop_front().unwrap();
            let re = r.pop_front().unwrap();
            //println!("? {:?}, {:?}", le, re);
            match (le, re) {
                (Packet::N(ln), Packet::L) => {
                    l.push_front(Packet::R);
                    l.push_front(Packet::N(ln));
                    l.push_front(Packet::L);
                    r.push_front(re);
                },
                (Packet::L, Packet::N(rn)) => {
                    r.push_front(Packet::R);
                    r.push_front(Packet::N(rn));
                    r.push_front(Packet::L);
                    l.push_front(le);
                },
                (Packet::N(ln), Packet::N(lr)) => {
                    if ln < lr {
                        return Ordering::Less;
                    } else if ln > lr {
                        return Ordering::Greater;
                    } else {
                        // continue
                    }
                },
                (Packet::L, Packet::L) => {
                    let c = cmp1(l, r);
                    if c != Ordering::Equal { return c; }
                },
                (Packet::R, Packet::R) => {
                    return Ordering::Equal;
                },
                (Packet::R, _) => {
                    return Ordering::Less;
                },
                (_, Packet::R) => {
                    return Ordering::Greater;
                },
            }
        }
    }

    let mut packets = Vec::new();
    for pair in pairs {
        packets.push(pair.l);
        packets.push(pair.r);
    }
    packets.push(VecDeque::from(vec![Packet::L, Packet::L, Packet::N(2), Packet::R, Packet::R]));
    packets.push(VecDeque::from(vec![Packet::L, Packet::L, Packet::N(6), Packet::R, Packet::R]));
    
    packets.sort_by(|a,b| cmp1(&mut a.clone(), &mut b.clone()));

    let mut index = 1;
    let mut out = 1;
    for ps in packets {
        if ps == VecDeque::from(vec![Packet::L, Packet::L, Packet::N(2), Packet::R, Packet::R])
            || ps == VecDeque::from(vec![Packet::L, Packet::L, Packet::N(6), Packet::R, Packet::R]) {
            out *= index;
            }
        index += 1;
        //println!("{:?}", ps);
    }
    println!("{out}");
}

fn main() {
    star1();
    star2();
}
