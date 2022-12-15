fn star1() {
    struct Hit { sx : i32, sy : i32, bx : i32, by : i32, r : i32 }

    fn dist(x0 : i32, y0 : i32, x1:i32, y1:i32) -> i32 {
        (x1-x0).abs() + (y1-y0).abs()
    }

    let mut hits = Vec::new();
    for line in std::fs::read_to_string("15.txt").unwrap().split("\n") {
        if line.len() == 0 { continue; }
        let split = line.split(" ").collect::<Vec<_>>();

        let sx = split[2][2..split[2].len()-1].parse::<i32>().unwrap();
        let sy = split[3][2..split[3].len()-1].parse::<i32>().unwrap();

        let bx = split[8][2..split[8].len()-1].parse::<i32>().unwrap();
        let by = split[9][2..].parse::<i32>().unwrap();

        let r = dist(sx,sy, bx,by);
        hits.push(Hit { sx, sy, bx, by, r });
        //println!("{sx},{sy} = {bx},{by}");
    }

    let mut sum = 0;
    //let y = 10;
    let y = 2000000;
    for x in -4000000..8000000 {
        let mut detectable = true;
        for hit in hits.iter() {
            if x == hit.bx && y == hit.by { detectable = true; break; }
            if dist(x,y, hit.sx,hit.sy) <= hit.r { detectable = false; break; }
        }
        //println!("{x},{y} = {detectable}");
        if !detectable { sum += 1; }
    }
    println!("{sum}");
    // 4415428 too low
}

fn star2() {
    struct Hit { sx : i32, sy : i32, bx : i32, by : i32, r : i32 }

    fn dist(x0 : i32, y0 : i32, x1:i32, y1:i32) -> i32 {
        (x1-x0).abs() + (y1-y0).abs()
    }

    let mut hits = Vec::new();
    for line in std::fs::read_to_string("15.txt").unwrap().split("\n") {
        if line.len() == 0 { continue; }
        let split = line.split(" ").collect::<Vec<_>>();

        let sx = split[2][2..split[2].len()-1].parse::<i32>().unwrap();
        let sy = split[3][2..split[3].len()-1].parse::<i32>().unwrap();

        let bx = split[8][2..split[8].len()-1].parse::<i32>().unwrap();
        let by = split[9][2..].parse::<i32>().unwrap();

        let r = dist(sx,sy, bx,by);
        hits.push(Hit { sx, sy, bx, by, r });
        //println!("{sx},{sy} = {bx},{by}");
    }

    //let max = 21;
    let max = 4000000 + 1;

    for y in 0..max {
        //println!("{y}");
        let mut nots = Vec::new();
        for hit in hits.iter() {
            let ywidth = hit.r - (y - hit.sy).abs();
            if ywidth < 0 { continue; }
            let x1 = hit.sx - ywidth;
            let x2 = hit.sx + ywidth;
            //println!("    {x1}..{x2}");
            nots.push((x1, x2));
        }
        nots.sort_by(|a,b| a.0.cmp(&b.0));
        //println!("## {nots:?}");
        let mut maxx = 0;
        for (x1, x2) in nots.iter() {
            let x1 = *x1;
            let x2 = *x2;
            if x1 > maxx {
                //println!("! y={y}, interval={x1},{x2}, maxx={maxx}");
                let mut signal : i64 = maxx as i64 *4000000 + y as i64;
                println!("{signal}");
                return;
            }
            if x2+1 > maxx { maxx = x2+1; }
        }
    }
}

fn main() {
    star1();
    star2();
}
