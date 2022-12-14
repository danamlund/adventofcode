fn star1() {
    use std::collections::HashSet;

    fn sgn(x : i32) -> i32 {
        if x < 0 { -1 } else if x > 0 { 1 } else { 0 }
    }
    
    let mut area = HashSet::new();
    let mut maxy = 0;
    for line in std::fs::read_to_string("14.txt").unwrap().split("\n") {
        if line.len() == 0 { continue; }
        let mut xy0 : Option<(i32,i32)> = None;
        for pos_str in line.split(" -> ") {
            let num_strs = pos_str.split(",").collect::<Vec<_>>();
            let xy : (i32, i32) = (num_strs[0].parse::<i32>().unwrap(), num_strs[1].parse::<i32>().unwrap());
            if let Some(xy0) = xy0 {
                let dx = sgn(xy.0 - xy0.0);
                let dy = sgn(xy.1 - xy0.1);
                for i in 0.. {
                    let xi = xy0.0 + dx*i;
                    let yi = xy0.1 + dy*i;
                    area.insert((xi, yi));
                    if yi > maxy { maxy = yi; }
                    //println!("{xy0:?} -> {xy:?}: {i}, {xi},{yi}");
                    if xi == xy.0 && yi == xy.1 { break; }
                }
            }
            xy0 = Some(xy);
        }
    }

    let mut sands = 0;
    loop {
        let mut tries = 0;
        let mut x = 500;
        let mut y = 0;
        loop {
            tries += 1;
            if y > maxy { tries = 1; break; }
            if !area.contains(&(x, y+1)) { y += 1; continue; }
            if !area.contains(&(x-1, y+1)) { x -= 1; y += 1; continue; }
            if !area.contains(&(x+1, y+1)) { x += 1; y += 1; continue; }
            break;
        }
        if tries == 1 { break; }
        area.insert((x,y)); 
        //println!("{sands}: {x},{y} {tries}");
        sands += 1;
    }

    println!("{sands}");
}

fn star2() {
    use std::collections::HashSet;

    fn sgn(x : i32) -> i32 {
        if x < 0 { -1 } else if x > 0 { 1 } else { 0 }
    }
    
    let mut area = HashSet::new();
    let mut maxy = 0;
    for line in std::fs::read_to_string("14.txt").unwrap().split("\n") {
        if line.len() == 0 { continue; }
        let mut xy0 : Option<(i32,i32)> = None;
        for pos_str in line.split(" -> ") {
            let num_strs = pos_str.split(",").collect::<Vec<_>>();
            let xy : (i32, i32) = (num_strs[0].parse::<i32>().unwrap(), num_strs[1].parse::<i32>().unwrap());
            if let Some(xy0) = xy0 {
                let dx = sgn(xy.0 - xy0.0);
                let dy = sgn(xy.1 - xy0.1);
                for i in 0.. {
                    let xi = xy0.0 + dx*i;
                    let yi = xy0.1 + dy*i;
                    area.insert((xi, yi));
                    if yi > maxy { maxy = yi; }
                    //println!("{xy0:?} -> {xy:?}: {i}, {xi},{yi}");
                    if xi == xy.0 && yi == xy.1 { break; }
                }
            }
            xy0 = Some(xy);
        }
    }

    fn print_sands(area : &HashSet<(i32,i32)>, minx : i32, maxx : i32, maxy : i32) {
        for y in 0..maxy+1 {
            for x in minx..maxx+1 {
                if area.contains(&(x,y)) { print!("O"); } else  { print!("."); }
            }
            println!();
        }
    }
    
    let mut sands = 0;
    loop {
        let mut tries = 0;
        let mut x = 500;
        let mut y = 0;
        loop {
            tries += 1;
            if y > maxy { break; }
            if !area.contains(&(x, y+1)) { y += 1; continue; }
            if !area.contains(&(x-1, y+1)) { x -= 1; y += 1; continue; }
            if !area.contains(&(x+1, y+1)) { x += 1; y += 1; continue; }
            break;
        }
        if tries == 1 { break; }
        area.insert((x,y)); 
        //println!("{sands}: {x},{y} {tries}");
        sands += 1;
    }
    sands += 1;

    //print_sands(&area, 500-20, 500+20, maxy+3);

    println!("{sands}");
    // 28190 too high
}


fn main() {
    star1();
    star2();
}
