// Change.Perform wants to mount a filesystem but the mount point isn't there.
func (s *changeSuite) TestPerformFilesystemMountWithoutMountPoint(c *check.C) {
	...
	s.sys.InsertFault(`lstat "/target"`, syscall.ENOENT)
	chg := &update.Change{Action: update.Mount,
		Entry: osutil.MountEntry{Name: "device", Dir: "/target", Type: "type"}}
	synth, err := chg.Perform(s.as)
	c.Assert(err, IsNil)
	c.Assert(synth, HasLen, 0)
	c.Assert(s.sys.RCalls(), testutil.SyscallsEqual, []testutil.CallResultError{
		{C: `lstat "/target"`, E: syscall.ENOENT},
		{C: `open "/" O_NOFOLLOW|O_CLOEXEC|O_DIRECTORY 0`, R: 3},
		{C: `mkdirat 3 "target" 0755`},
		{C: `openat 3 "target" O_NOFOLLOW|O_CLOEXEC|O_DIRECTORY 0`, R: 4},
		{C: `fchown 4 0 0`},
		{C: `close 4`},
		{C: `close 3`},
		{C: `mount "device" "/target" "type" 0 ""`},
	})
}
