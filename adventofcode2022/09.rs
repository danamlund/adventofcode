use std::collections::HashSet;

fn p(visited : &HashSet<(i32,i32)>) {
    let mut minx = 99999999;
    let mut maxx = -9999999;
    let mut miny = 99999999;
    let mut maxy = -9999999;
    for (x,y) in visited {
        if *x < minx { minx = *x; }
        if *x > maxx { maxx = *x; }
        if *y < miny { miny = *y; }
        if *y > maxy { maxy = *y; }
    }
    minx -= 1; maxx += 1;
    for y in miny..maxy+1 {
        for x in minx..maxx+1 {
            print!("{}", if visited.contains(&(x,y)) { "#" } else { "." });
        }
        println!();
    }
    println!();
}

fn sign(x : i32) -> i32 {
    if x < 0 { -1 } else if x > 0 { 1 } else { 0 }
}

fn star1() {
    let mut visited = HashSet::new();
    let (mut hx, mut hy) = (0,0);
    let (mut tx, mut ty) = (0,0);
    visited.insert((tx,ty));
    for line in std::fs::read_to_string("09.txt").unwrap().split("\n") {
        if line.len() == 0 { continue; }
        let dir = line.chars().nth(0).unwrap();
        let amount = line[2..].parse::<i32>().unwrap();
        //println!("## {dir} {amount}");
        for _ in 0..amount {
            if dir == 'R' { hx += 1; }
            if dir == 'L' { hx -= 1; }
            if dir == 'U' { hy -= 1; }
            if dir == 'D' { hy += 1; }
            //println!("1  {hx},{hy} -> {tx},{ty}");
            if i32::abs(hx - tx) > 1 || i32::abs(hy - ty) > 1 {
                tx += sign(hx - tx);
                ty += sign(hy - ty);
            }

            let mut v = HashSet::new();
            v.insert((tx,ty));
            let tx0 = tx;
            let ty0 = ty;
            
            // if ty == hy && tx + 2 == hx { tx += 1; }
            // else if ty == hy && tx - 2 == hx { tx -= 1; }
            // else if tx == hx && ty + 2 == hy { ty += 1; }
            // else if tx == hx && ty - 2 == hy { ty -= 1; }
            // else if tx + 2 == hx || tx - 2 == hx || ty + 2 == hy || ty - 2 == hy {
            //     tx += sign(hx - tx);
            //     ty += sign(hy - ty);
            //     //println!("{hx},{hy}: {tx0},{ty0} -> {tx},{ty}");
            // }

            // if tx0 != tx || ty0 != ty {
            //     println!("{hx},{hy}: {tx0},{ty0} -> {tx},{ty}");
            //     v.insert((tx,ty));
            //     p(&v);
            // }

            //  TTT
            // T   T
            // T H T
            // T   T
            //  TTT
            

            //println!("2  {hx},{hy} -> {tx},{ty}");
            visited.insert((tx, ty));
            //p(&visited);
            //println!();
        }
    }

    //println!("{:?}", visited);
    //p(&visited);

    println!("{}", visited.len());
    // 3001 too low
}

fn star2() {
    let mut visited = HashSet::new();
    let mut knots = vec![(0,0); 10];
    visited.insert(knots[9]);
    for line in std::fs::read_to_string("09.txt").unwrap().split("\n") {
        if line.len() == 0 { continue; }
        let dir = line.chars().nth(0).unwrap();
        let amount = line[2..].parse::<i32>().unwrap();
        //println!("## {dir} {amount}");
        for _ in 0..amount {
            let (mut hx,mut hy) = knots[0];
            if dir == 'R' { hx += 1; }
            if dir == 'L' { hx -= 1; }
            if dir == 'U' { hy -= 1; }
            if dir == 'D' { hy += 1; }
            knots[0] = (hx,hy);

            for i in 1..knots.len() {
                let (hx,hy) = knots[i-1];
                let (mut tx,mut ty) = knots[i];
                if i32::abs(hx - tx) > 1 || i32::abs(hy - ty) > 1 {
                    tx += sign(hx - tx);
                    ty += sign(hy - ty);
                }
                knots[i] = (tx,ty);
            }

            visited.insert(knots[9]);
        }
    }

    //println!("{:?}", visited);
    //p(&visited);

    println!("{}", visited.len());
    // 3001 too low
}


fn main() {
    star1();
    star2();
}
