package mock_exec

import (
	"errors"
	"io"
	"testing"

	"github.com/pdutton/go-mocks/internal/testutil"
	"go.uber.org/mock/gomock"
)

// TestMockCmd_Run tests running a command.
func TestMockCmd_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	mockCmd.EXPECT().Run().Return(nil)

	err := mockCmd.Run()
	testutil.AssertNil(t, err)
}

// TestMockCmd_RunError tests command execution error.
func TestMockCmd_RunError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	expectedErr := errors.New("exit status 1")
	mockCmd.EXPECT().Run().Return(expectedErr)

	err := mockCmd.Run()
	testutil.AssertError(t, expectedErr, err)
}

// TestMockCmd_Start tests starting a command.
func TestMockCmd_Start(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	mockCmd.EXPECT().Start().Return(nil)

	err := mockCmd.Start()
	testutil.AssertNil(t, err)
}

// TestMockCmd_Wait tests waiting for command completion.
func TestMockCmd_Wait(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	mockCmd.EXPECT().Wait().Return(nil)

	err := mockCmd.Wait()
	testutil.AssertNil(t, err)
}

// TestMockCmd_StartWait tests Start then Wait sequence.
func TestMockCmd_StartWait(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	gomock.InOrder(
		mockCmd.EXPECT().Start().Return(nil),
		mockCmd.EXPECT().Wait().Return(nil),
	)

	err1 := mockCmd.Start()
	testutil.AssertNil(t, err1)

	err2 := mockCmd.Wait()
	testutil.AssertNil(t, err2)
}

// TestMockCmd_Output tests capturing command output.
func TestMockCmd_Output(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	expectedOutput := []byte("hello world")
	mockCmd.EXPECT().Output().Return(expectedOutput, nil)

	output, err := mockCmd.Output()
	testutil.AssertNil(t, err)
	testutil.AssertBytes(t, expectedOutput, output)
}

// TestMockCmd_OutputError tests Output error handling.
func TestMockCmd_OutputError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	expectedErr := errors.New("command failed")
	mockCmd.EXPECT().Output().Return([]byte(nil), expectedErr)

	output, err := mockCmd.Output()
	if len(output) != 0 {
		t.Errorf("expected empty output, got %v", output)
	}
	testutil.AssertError(t, expectedErr, err)
}

// TestMockCmd_CombinedOutput tests capturing stdout and stderr.
func TestMockCmd_CombinedOutput(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	expectedOutput := []byte("combined output")
	mockCmd.EXPECT().CombinedOutput().Return(expectedOutput, nil)

	output, err := mockCmd.CombinedOutput()
	testutil.AssertNil(t, err)
	testutil.AssertBytes(t, expectedOutput, output)
}

// TestMockCmd_CombinedOutputError tests CombinedOutput error.
func TestMockCmd_CombinedOutputError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	expectedErr := errors.New("command failed")
	expectedOutput := []byte("error message")
	mockCmd.EXPECT().CombinedOutput().Return(expectedOutput, expectedErr)

	output, err := mockCmd.CombinedOutput()
	testutil.AssertBytes(t, expectedOutput, output)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockCmd_StdinPipe tests getting stdin pipe.
func TestMockCmd_StdinPipe(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	var expectedPipe io.WriteCloser
	mockCmd.EXPECT().StdinPipe().Return(expectedPipe, nil)

	pipe, err := mockCmd.StdinPipe()
	testutil.AssertNil(t, pipe)
	testutil.AssertNil(t, err)
}

// TestMockCmd_StdoutPipe tests getting stdout pipe.
func TestMockCmd_StdoutPipe(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	var expectedPipe io.ReadCloser
	mockCmd.EXPECT().StdoutPipe().Return(expectedPipe, nil)

	pipe, err := mockCmd.StdoutPipe()
	testutil.AssertNil(t, pipe)
	testutil.AssertNil(t, err)
}

// TestMockCmd_StderrPipe tests getting stderr pipe.
func TestMockCmd_StderrPipe(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	var expectedPipe io.ReadCloser
	mockCmd.EXPECT().StderrPipe().Return(expectedPipe, nil)

	pipe, err := mockCmd.StderrPipe()
	testutil.AssertNil(t, pipe)
	testutil.AssertNil(t, err)
}

// TestMockCmd_Path tests getting command path.
func TestMockCmd_Path(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	mockCmd.EXPECT().Path().Return("/bin/ls")

	path := mockCmd.Path()
	testutil.AssertEqual(t, "/bin/ls", path)
}

// TestMockCmd_Args tests getting command arguments.
func TestMockCmd_Args(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	args := []string{"ls", "-la", "/tmp"}
	mockCmd.EXPECT().Args().Return(args)

	result := mockCmd.Args()
	testutil.AssertEqual(t, 3, len(result))
	testutil.AssertEqual(t, "ls", result[0])
}

// TestMockCmd_Dir tests getting working directory.
func TestMockCmd_Dir(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	mockCmd.EXPECT().Dir().Return("/home/user")

	dir := mockCmd.Dir()
	testutil.AssertEqual(t, "/home/user", dir)
}

// TestMockCmd_Env tests getting environment variables.
func TestMockCmd_Env(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	env := []string{"PATH=/usr/bin", "HOME=/home/user"}
	mockCmd.EXPECT().Env().Return(env)

	result := mockCmd.Env()
	testutil.AssertEqual(t, 2, len(result))
	testutil.AssertEqual(t, "PATH=/usr/bin", result[0])
}

// TestMockCmd_Environ tests getting full environment.
func TestMockCmd_Environ(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	environ := []string{"PATH=/usr/bin", "HOME=/home/user", "USER=test"}
	mockCmd.EXPECT().Environ().Return(environ)

	result := mockCmd.Environ()
	testutil.AssertEqual(t, 3, len(result))
}

// TestMockCmd_Stdin tests getting stdin.
func TestMockCmd_Stdin(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	var expectedStdin io.Reader
	mockCmd.EXPECT().Stdin().Return(expectedStdin)

	stdin := mockCmd.Stdin()
	testutil.AssertNil(t, stdin)
}

// TestMockCmd_Stdout tests getting stdout.
func TestMockCmd_Stdout(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	var expectedStdout io.Writer
	mockCmd.EXPECT().Stdout().Return(expectedStdout)

	stdout := mockCmd.Stdout()
	testutil.AssertNil(t, stdout)
}

// TestMockCmd_Stderr tests getting stderr.
func TestMockCmd_Stderr(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	var expectedStderr io.Writer
	mockCmd.EXPECT().Stderr().Return(expectedStderr)

	stderr := mockCmd.Stderr()
	testutil.AssertNil(t, stderr)
}

// TestMockCmd_String tests string representation.
func TestMockCmd_String(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd := NewMockCmd(ctrl)

	mockCmd.EXPECT().String().Return("ls -la /tmp")

	str := mockCmd.String()
	testutil.AssertEqual(t, "ls -la /tmp", str)
}

// TestMockExec_NewCommand tests creating a new command.
func TestMockExec_NewCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockExec := NewMockExec(ctrl)
	mockCmd := NewMockCmd(ctrl)

	mockExec.EXPECT().NewCommand("ls").Return(mockCmd)

	cmd := mockExec.NewCommand("ls")
	testutil.AssertNotNil(t, cmd)
}

// TestMockExec_LookPath tests looking up executable path.
func TestMockExec_LookPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockExec := NewMockExec(ctrl)

	mockExec.EXPECT().LookPath("ls").Return("/bin/ls", nil)

	path, err := mockExec.LookPath("ls")
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, "/bin/ls", path)
}

// TestMockExec_LookPathError tests LookPath error handling.
func TestMockExec_LookPathError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockExec := NewMockExec(ctrl)

	expectedErr := errors.New("executable not found")
	mockExec.EXPECT().LookPath("nonexistent").Return("", expectedErr)

	path, err := mockExec.LookPath("nonexistent")
	testutil.AssertEqual(t, "", path)
	testutil.AssertError(t, expectedErr, err)
}

// TestMockCmd_PipelineExecution tests command pipeline.
func TestMockCmd_PipelineExecution(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCmd1 := NewMockCmd(ctrl)
	mockCmd2 := NewMockCmd(ctrl)

	var pipeWriter io.WriteCloser
	var pipeReader io.ReadCloser

	gomock.InOrder(
		mockCmd1.EXPECT().StdoutPipe().Return(pipeReader, nil),
		mockCmd2.EXPECT().StdinPipe().Return(pipeWriter, nil),
		mockCmd1.EXPECT().Start().Return(nil),
		mockCmd2.EXPECT().Start().Return(nil),
		mockCmd1.EXPECT().Wait().Return(nil),
		mockCmd2.EXPECT().Wait().Return(nil),
	)

	// Simulate pipeline: cmd1 | cmd2
	_, err1 := mockCmd1.StdoutPipe()
	testutil.AssertNil(t, err1)

	_, err2 := mockCmd2.StdinPipe()
	testutil.AssertNil(t, err2)

	err3 := mockCmd1.Start()
	testutil.AssertNil(t, err3)

	err4 := mockCmd2.Start()
	testutil.AssertNil(t, err4)

	err5 := mockCmd1.Wait()
	testutil.AssertNil(t, err5)

	err6 := mockCmd2.Wait()
	testutil.AssertNil(t, err6)
}
