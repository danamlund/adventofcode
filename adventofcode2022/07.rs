use std::collections::HashMap;

struct F {
    file_to_size : HashMap<String, i32>
}

impl F {
    fn new() -> F {
        F { file_to_size: HashMap::new() }
    }

    fn set(&mut self, f : &str, size :i32) {
        self.file_to_size.insert(f.to_string(), size);
    }

    fn dirsizes(&self) -> HashMap<String, i32> {
        let mut dir_to_size = HashMap::new();
        for (path, size) in self.file_to_size.iter() {
            if *size == 0 { continue; }
            let mut path0 = path.to_string();
            //println!("##1 {path0}");
            while path0.len() > 1 {
                while path0.pop().unwrap() != '/' {}
                if path0 == "" { path0.push('/'); }
                //println!("## {path0} = {size}");
                dir_to_size.insert(path0.clone(), *size + dir_to_size.get(&path0).unwrap_or(&0));
            }
        }
        return dir_to_size;
    }
}

fn star1() {
    let txt = std::fs::read_to_string("07.txt").unwrap();
    let mut pwd = String::new();
    let mut f = F::new();
    f.set("/", 0);
    for line in txt.split("\n") {
        if line == "" { continue; }
        //println!("{line}   - {pwd}");
        if line.starts_with("$ cd ") {
            let cd = &line["$ cd ".len()..];
            if cd == "/" {
                pwd.clear();
                pwd.push_str("/");
            } else if cd == ".." {
                while pwd.pop().unwrap() != '/' {}
                if pwd == "" { pwd.push_str("/"); }
            } else {
                if pwd != "/" { pwd.push_str("/"); }
                pwd.push_str(cd);
            }
        } else if line == "$ ls" {
            // skip
        } else {
            let (size_str, file) = line.split_once(' ').unwrap();
            let size = size_str.parse().unwrap_or(0);
            
            let mut path = String::new();
            path.push_str(&pwd);
            if path != "/" { path.push_str("/"); }
            path.push_str(file);
            f.set(&path, size);
        }
    }

    //println!("{:?}", f.file_to_size);

    let ds = f.dirsizes();
    //println!("{:?}", ds);

    let mut sum = 0;
    for (_, size) in ds {
        if size <= 100000 {
            sum += size;
        }
    }
    println!("{sum}");
}


fn star2() {
    let txt = std::fs::read_to_string("07.txt").unwrap();
    let mut pwd = String::new();
    let mut f = F::new();
    f.set("/", 0);
    for line in txt.split("\n") {
        if line == "" { continue; }
        //println!("{line}   - {pwd}");
        if line.starts_with("$ cd ") {
            let cd = &line["$ cd ".len()..];
            if cd == "/" {
                pwd.clear();
                pwd.push_str("/");
            } else if cd == ".." {
                while pwd.pop().unwrap() != '/' {}
                if pwd == "" { pwd.push_str("/"); }
            } else {
                if pwd != "/" { pwd.push_str("/"); }
                pwd.push_str(cd);
            }
        } else if line == "$ ls" {
            // skip
        } else {
            let (size_str, file) = line.split_once(' ').unwrap();
            let size = size_str.parse().unwrap_or(0);
            
            let mut path = String::new();
            path.push_str(&pwd);
            if path != "/" { path.push_str("/"); }
            path.push_str(file);
            f.set(&path, size);
        }
    }

    //println!("{:?}", f.file_to_size);

    let ds = f.dirsizes();
    //println!("{:?}", ds);

    let to_free = 30_000_000 - (70_000_000 - ds["/"]);
    
    let mut min = 99999999;
    for (_, size) in ds {
        if size < min && size >= to_free {
            min = size;
        }
    }
    println!("{min}");
}


fn main() {
    star1();
    star2();
}
