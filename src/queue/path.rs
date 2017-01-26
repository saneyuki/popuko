use std::path::{Component, Path, PathBuf};

pub fn create_safe_path<P>(root: P, path: P) -> Option<PathBuf>
    where P: AsRef<Path> + ::std::fmt::Display
{
    let mut p = PathBuf::new();
    p.push(format!("{}", root));
    p.push(format!("./{}", path));

    for c in p.components() {
        match c {
            Component::ParentDir |
            Component::CurDir => {
                return None;
            }
            _ => (),
        }
    }

    Some(p)
}

#[cfg(test)]
mod tests {
    use super::*;

    macro_rules! test_create_safe_path {
        ( $($name:ident: ($input_root:expr, $input_path:expr, $expected:expr))* ) => {
            $(
                #[test]
                fn $name() {
                    let r = create_safe_path($input_root, $input_path);
                    assert_eq!(r, $expected);
                }
            )*
        }
    }

    test_create_safe_path! {
        test_create_safe_path1: ( "/a/b/c", "d/e", Some(PathBuf::from("/a/b/c/d/e")) )
        test_create_safe_path2: ( "/a/b/c", "./d/e", Some(PathBuf::from("/a/b/c/d/e")) )
        test_create_safe_path3: ( "/a/b/c", "/d/e", Some(PathBuf::from("/a/b/c/d/e")) )
        test_create_safe_path4: ( "a/b/c", "/d/e", Some(PathBuf::from("a/b/c/d/e")) )
        test_create_safe_path5: ( "/a/b/..", "/d/e", None )
        test_create_safe_path6: ( "/a/b/c", "../d/e", None )
        test_create_safe_path7: ( "/a/b/c", "../../e", None )
        test_create_safe_path8: ( "/a/b/c", "../~/e", None )
        test_create_safe_path9: ( "C:\\server\\share", "../~/e", None )
        test_create_safe_path10: ( "\\\\server\\share", "../~/e", None )
        //test_create_safe_path11: ( "/a/b/c", "C:\\server\\share", None )
        //test_create_safe_path12: ( "/a/b/c", "\\\\server\\share", None )
    }
}
