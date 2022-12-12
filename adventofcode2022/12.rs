
fn star1() {
    use std::collections::HashMap;
    use std::collections::VecDeque;
    let mut heights = HashMap::new();
    let mut y = 0;
    let mut start = (-1,-1);
    let mut end = (-1,-1);
    for line in std::fs::read_to_string("12.txt").unwrap().split("\n") {
        let mut x = 0;
        for c in line.chars() {
            let mut h = c as i32 - 'a' as i32;
            if c == 'S' { start = (x,y); h = 0; }
            if c == 'E' { end = (x,y); h = 'z' as i32 - 'a' as i32; }
            heights.insert((x,y), h);
            x += 1;
        }
        y += 1;
    }

    let mut dists = HashMap::new();
    dists.insert(start, 0);
    let mut cloud = VecDeque::new();
    cloud.push_back(start);

    fn print(heights : &HashMap<(i32,i32), i32>, dists : &HashMap<(i32,i32), i32>) {
        let mut minx = 99999;
        let mut maxx = -9999;
        let mut miny = 99999;
        let mut maxy = -9999;
        for (x,y) in heights.keys() {
            if *x < minx { minx = *x; }
            if *x > maxx { maxx = *x; }
            if *y < miny { miny = *y; }
            if *y > maxy { maxy = *y; }
        }
        for y in miny..maxy+1 {
            for x in  minx..maxx+1 {
                print!("{:2} ", dists.get(&(x,y)).unwrap_or(&-1));
            }
            println!();
        }
    }
    
    while cloud.len() > 0 {
        cloud.make_contiguous().sort_by(|a,b| dists.get(&a).unwrap_or(&99999).cmp(dists.get(&b).unwrap_or(&99999)));
        //println!("{:?}", cloud);
        let cur = cloud.pop_front().unwrap();
        let cur_height = *heights.get(&cur).unwrap();
        let cur_dist = *dists.get(&cur).unwrap();
        //println!("{:?} h={cur_height} d={}", cur, dists.get(&cur).unwrap());
        for d in vec![(-1,0), (1,0), (0,-1), (0,1)] {
            let neighbor = (cur.0+d.0, cur.1+d.1);
            if !heights.contains_key(&neighbor) { continue; }
            if dists.contains_key(&neighbor) { continue; }
            let neighbor_height = *heights.get(&neighbor).unwrap();
            //println!("  ? {:?} h={neighbor_height}", neighbor);
            if cur_height >= 0 && neighbor_height >= 0 && neighbor_height > cur_height + 1 { continue; }
            if !dists.contains_key(&neighbor) || cur_dist + 1 < *dists.get(&neighbor).unwrap() {
                dists.insert(neighbor, cur_dist + 1);
            }
            cloud.push_back(neighbor);
        }
        //print(&heights, &dists);
    }
    print(&heights, &dists);
    
    println!("{}", dists.get(&end).unwrap());
    // 519 too high
}

fn star2() {
    use std::collections::HashMap;
    use std::collections::VecDeque;
    let mut heights = HashMap::new();
    let mut y = 0;
    let mut end = (-1,-1);
    for line in std::fs::read_to_string("12.txt").unwrap().split("\n") {
        let mut x = 0;
        for c in line.chars() {
            let mut h = c as i32 - 'a' as i32;
            if c == 'S' { h = 0; }
            if c == 'E' { end = (x,y); h = 'z' as i32 - 'a' as i32; }
            heights.insert((x,y), h);
            x += 1;
        }
        y += 1;
    }

    let mut dists = HashMap::new();
    let mut cloud = VecDeque::new();

    fn print(heights : &HashMap<(i32,i32), i32>, dists : &HashMap<(i32,i32), i32>) {
        let mut minx = 99999;
        let mut maxx = -9999;
        let mut miny = 99999;
        let mut maxy = -9999;
        for (x,y) in heights.keys() {
            if *x < minx { minx = *x; }
            if *x > maxx { maxx = *x; }
            if *y < miny { miny = *y; }
            if *y > maxy { maxy = *y; }
        }
        for y in miny..maxy+1 {
            for x in  minx..maxx+1 {
                print!("{:2} ", dists.get(&(x,y)).unwrap_or(&-1));
            }
            println!();
        }
    }

    let mut mindist = 9999;
    for start in heights.keys() {
        if heights.get(start).unwrap() != &0 { continue; }
        //println!("{:?}", start2);
        dists.clear();
        cloud.clear();
        dists.insert(*start, 0);
        cloud.push_back(*start);
        while cloud.len() > 0 {
            cloud.make_contiguous().sort_by(|a,b| dists.get(&a).unwrap_or(&99999).cmp(dists.get(&b).unwrap_or(&99999)));
            //println!("{:?}", cloud);
            let cur = cloud.pop_front().unwrap();
            let cur_height = *heights.get(&cur).unwrap();
            let cur_dist = *dists.get(&cur).unwrap();
            //println!("{:?} h={cur_height} d={}", cur, dists.get(&cur).unwrap());
            for d in vec![(-1,0), (1,0), (0,-1), (0,1)] {
                let neighbor = (cur.0+d.0, cur.1+d.1);
                if !heights.contains_key(&neighbor) { continue; }
                if dists.contains_key(&neighbor) { continue; }
                let neighbor_height = *heights.get(&neighbor).unwrap();
                //println!("  ? {:?} h={neighbor_height}", neighbor);
                if cur_height >= 0 && neighbor_height >= 0 && neighbor_height > cur_height + 1 { continue; }
                if !dists.contains_key(&neighbor) || cur_dist + 1 < *dists.get(&neighbor).unwrap() {
                    dists.insert(neighbor, cur_dist + 1);
                }
                cloud.push_back(neighbor);
            }
            //print(&heights, &dists);
        }
        //print(&heights, &dists);
        let dist = dists.get(&end).unwrap_or(&9999);
        if *dist < mindist {
            mindist = *dist;
            //println!("{}", mindist);
            //print(&heights, &dists);
        }
    }
    
    println!("{}", mindist);
    // 519 too high
}

fn main() {
    star1();
    star2();
}
