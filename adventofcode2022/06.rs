fn star1() {
    let txt = std::fs::read_to_string("06.txt").unwrap();
    let mut prev = [0 as char; 4];
    let mut previ = 0;
    for (i, c) in txt.chars().enumerate() {
        prev[previ] = c;
        previ = (previ+1) % prev.len();
        //println!("{i} = {c}");
        if i > 3 {
            let mut has_same = false;
            for j in 0..prev.len() {
                for k in j+1..prev.len() {
                    if prev[j] == prev[k] {
                        has_same = true;
                    }
                }
            }
            if !has_same {
                println!("{}", i+1);
                break;
            }
        }
    }
    // 1274 too low
}

fn star2() {
    let txt = std::fs::read_to_string("06.txt").unwrap();
    let mut prev = [0 as char; 14];
    let mut previ = 0;
    for (i, c) in txt.chars().enumerate() {
        prev[previ] = c;
        previ = (previ+1) % prev.len();
        //println!("{i} = {c}");
        if i > 3 {
            let mut has_same = false;
            for j in 0..prev.len() {
                for k in j+1..prev.len() {
                    if prev[j] == prev[k] {
                        has_same = true;
                    }
                }
            }
            if !has_same {
                println!("{}", i+1);
                break;
            }
        }
    }
}

fn main() {
    star1();
    star2();
}
