use std::collections::BTreeMap;
use std::fs;
use std::io;
use std::path::{Path, PathBuf};
use std::sync::{Arc, RwLock};
use std::sync::mpsc::{Receiver, Sender};
use std::thread;

use queue::handle::QueueHandle;

pub enum FileQueryInput {
    Exit,
}

pub struct FileRepository {
    root: PathBuf,
    map: BTreeMap<String, Arc<RwLock<()>>>,

    receiver: Receiver<FileQueryInput>,
}

const QUEUE_REPO_DIR: &'static str = "./queue";

impl FileRepository {

    fn new(fileroot: &Path, recv: Receiver<FileQueryInput>) -> FileRepository {
        let q = Path::new(QUEUE_REPO_DIR);
        let root = fileroot.join(q);
        let map = BTreeMap::new();
        FileRepository {
            root: root,
            map: map,
            receiver: recv,
        }
    }

    fn init(&self) -> io::Result<()> {
        fs::create_dir(&self.root)
    }

    fn root(&self) -> &Path {
        self.root.as_path()
    }

    fn load(&mut self, owner: &str, name: &str, sender: Sender<Option<QueueHandle>>) {
        let repo = String::from(format!("{}/{}", owner, name));
        let filelock = self.map.entry(repo.clone()).or_insert_with( || {
            let lock = RwLock::new(());
            Arc::new(lock)
        });

        let root = self.root.clone();
        let filelock = filelock.clone();
        thread::spawn(move || {
            if let Err(e) = filelock.read() {
                let msg = format!("{}", e);
                unreachable!(msg);
            }

            // TODO: create safe path.
            let filepath = root.join(repo);
        });
    }
}

pub struct MergeQueueFile {
    version: u32,
}

#[cfg(test)]
mod tests {
    use std::path::Path;
    use std::sync::mpsc::channel;

    use super::*;

    macro_rules! test_file_repository_root_test {
        ( $($name:ident: $input:expr, $expected:expr,)* ) => {
            $(
                #[test]
                fn $name() {
                    let r = Path::new($input);
                    let (_, rx) = channel();
                    let repo = FileRepository::new(&r, rx);
                    assert_eq!(repo.root(), Path::new($expected));
                }
            )*
        }
    }

    test_file_repository_root_test! {
        test_file_repository_root1: "/a/b", "/a/b/queue",
        test_file_repository_root2: "/a/b/", "/a/b/queue",
    }
}
