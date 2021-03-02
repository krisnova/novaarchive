# Metrics

We look at `metrics` in this library a set of key/value pairs.
Where a `key` is a unique identifier for a `record` and the `value` is what the `record` contains.

We store meaningful metrics in various backends that can be used later to query for them.

### Supported storage 

 - File System / Disk 
 - ZFS

### Conventions 

The `Get()` and `Set()` methods are thread safe and atomic with regard to persistence to disk.
