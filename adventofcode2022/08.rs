use std::collections::HashSet;

#[derive(Debug)]
struct Forest {
    forest : Vec<i8>,
    width : i32,
}

impl Forest {
    fn new(width : i32) -> Forest {
        return Forest { forest: Vec::new(), width };
    }

    fn i(&self, x:i32, y:i32) -> usize {
        return (y*self.width + x) as usize;
    }
    fn set(&mut self, x : i32, y : i32, v : i8) {
        let i = self.i(x,y);
        while i >= self.forest.len() { self.forest.push(-1); }
        self.forest[i] = v;
    }
    fn get(&self, x:i32, y:i32) -> i8 {
        return self.forest[self.i(x,y)];
    }
}

fn star1() {
    let mut f = Forest::new(0);
    let mut y = 0;

    for line in std::fs::read_to_string("08.txt").unwrap().split("\n") {
        if f.width == 0 { f.width = line.len() as i32; }
        for x in 0..line.len() as i32 {
            f.set(x, y, line.chars().nth(x as usize).unwrap() as i8 - '0' as i8);
        }
        y += 1;
    }

    //println!("{:?}", f);
    //println!("last={}", f.get(f.width-1, f.width-1));

// 01234
//
// 30373 0
// 25512 1
// 65332 2
// 33549 3
// 35390 4

    let mut visible : HashSet<(i32,i32)> = HashSet::new();
    for y in 0..f.width {
        visible.insert((0,y));
        visible.insert((f.width-1,y));
        let mut blockh = 0;
        for x in 0..f.width {
            let h = f.get(x,y);
            if h > blockh {
                //println!("##1 {x},{y}: {h}>{blockh}");
                visible.insert((x,y));
                blockh = h;
            }
        }
        blockh = 0;
        for x in (0..f.width).rev() {
            let h = f.get(x,y);
            if h > blockh {
                //println!("##2 {x},{y}: {h}>{blockh}");
                visible.insert((x,y));
                blockh = h;
            }
        }
    }
    for x in 0..f.width {
        visible.insert((x,0));
        visible.insert((x,f.width-1));
        let mut blockh = 0;
        for y in 0..f.width {
            let h = f.get(x,y);
            if h > blockh {
                //println!("##3 {x},{y}: {h}>{blockh}");
                visible.insert((x,y));
                blockh = h;
            }
        }
        blockh = 0;
        for y in (0..f.width).rev() {
            let h = f.get(x,y);
            if h > blockh {
                //println!("##4 {x},{y}: {h}>{blockh}");
                visible.insert((x,y));
                blockh = h;
            }
        }
    }

    //println!("{:?}", visible);
    //println!("{:?}", visible.iter().filter(|(x,y)| *x != 0 && *y != 0).cloned().collect::<Vec<(i32,i32)>>());
    //println!("{:?}", visible.iter().filter(|(x,y)| *x == 0 || *y == 0).cloned().collect::<Vec<(i32,i32)>>());

    println!("{}", visible.len());
}


fn star2() {
    let mut f = Forest::new(0);
    let mut y = 0;

    for line in std::fs::read_to_string("08.txt").unwrap().split("\n") {
        if f.width == 0 { f.width = line.len() as i32; }
        for x in 0..line.len() as i32 {
            f.set(x, y, line.chars().nth(x as usize).unwrap() as i8 - '0' as i8);
        }
        y += 1;
    }

    let mut maxscore = 0;
    for y in 0..f.width {
        for x in 0..f.width {
            let mut score = 1;
            for (dx,dy) in vec![(0,-1),(0,1),(-1,0),(1,0)] {
                let mut trees = 0;
                let h = f.get(x,y);
                for i in 1..f.width {
                    let (x0,y0) = (x+dx*i, y+dy*i);
                    if x0 < 0 || x0 >= f.width { break; }
                    if y0 < 0 || y0 >= f.width { break; }
                    
                    let h0 = f.get(x0,y0);
                    trees += 1;
                    if h0 >= h { break; }
                }
                score *= trees;
                //println!("## {x},{y}, {dx},{dy} = {trees}");
            }
            //println!("## {x},{y} = {score}");
            if score > maxscore { maxscore = score; }
        }
    }
    println!("{}", maxscore);
}

fn main() {
    star1();
    star2();
}
