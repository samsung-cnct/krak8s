// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "krak8s": CLI Commands
//
// Command:
// $ goagen
// --design=krak8s/design
// --out=$(GOPATH)/src/krak8s
// --version=v1.2.0-dirty

package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"krak8s/client"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// CreateApplicationCommand is the command line data structure for the create action of application
	CreateApplicationCommand struct {
		Payload     string
		ContentType string
		// namespace identifier
		Ns string
		// project name
		Project     string
		PrettyPrint bool
	}

	// DeleteApplicationCommand is the command line data structure for the delete action of application
	DeleteApplicationCommand struct {
		Chart string
		// namespace identifier
		Ns string
		// project name
		Project     string
		PrettyPrint bool
	}

	// GetApplicationCommand is the command line data structure for the get action of application
	GetApplicationCommand struct {
		Chart string
		// namespace identifier
		Ns string
		// project name
		Project     string
		PrettyPrint bool
	}

	// ListApplicationCommand is the command line data structure for the list action of application
	ListApplicationCommand struct {
		// namespace identifier
		Ns string
		// project name
		Project     string
		PrettyPrint bool
	}

	// CreateClusterCommand is the command line data structure for the create action of cluster
	CreateClusterCommand struct {
		Payload     string
		ContentType string
		// namespace identifier
		Ns string
		// project name
		Project     string
		PrettyPrint bool
	}

	// DeleteClusterCommand is the command line data structure for the delete action of cluster
	DeleteClusterCommand struct {
		// namespace identifier
		Ns string
		// project name
		Project     string
		PrettyPrint bool
	}

	// GetClusterCommand is the command line data structure for the get action of cluster
	GetClusterCommand struct {
		// namespace identifier
		Ns string
		// project name
		Project     string
		PrettyPrint bool
	}

	// CreateNamespaceCommand is the command line data structure for the create action of namespace
	CreateNamespaceCommand struct {
		Payload     string
		ContentType string
		// project name
		Project     string
		PrettyPrint bool
	}

	// DeleteNamespaceCommand is the command line data structure for the delete action of namespace
	DeleteNamespaceCommand struct {
		// namespace identifier
		Ns string
		// project name
		Project     string
		PrettyPrint bool
	}

	// GetNamespaceCommand is the command line data structure for the get action of namespace
	GetNamespaceCommand struct {
		// namespace identifier
		Ns string
		// project name
		Project     string
		PrettyPrint bool
	}

	// ListNamespaceCommand is the command line data structure for the list action of namespace
	ListNamespaceCommand struct {
		// project name
		Project     string
		PrettyPrint bool
	}

	// CreateProjectCommand is the command line data structure for the create action of project
	CreateProjectCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// DeleteProjectCommand is the command line data structure for the delete action of project
	DeleteProjectCommand struct {
		// project name
		Project     string
		PrettyPrint bool
	}

	// GetProjectCommand is the command line data structure for the get action of project
	GetProjectCommand struct {
		// project name
		Project     string
		PrettyPrint bool
	}

	// ListProjectCommand is the command line data structure for the list action of project
	ListProjectCommand struct {
		PrettyPrint bool
	}

	// DownloadCommand is the command line data structure for the download command.
	DownloadCommand struct {
		// OutFile is the path to the download output file.
		OutFile string
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: `create action`,
	}
	tmp1 := new(CreateApplicationCommand)
	sub = &cobra.Command{
		Use:   `application ["/v1/projects/PROJECT/ns/NS/app"]`,
		Short: `Manage {create, delete}, and get namespaces's Application(s)`,
		Long: `Manage {create, delete}, and get namespaces's Application(s)

Payload example:

{
   "name": "Iste voluptas debitis voluptatem illum.",
   "registry": "Corrupti omnis atque maxime autem.",
   "set": "Ea corporis eaque id saepe aut provident.",
   "version": "Aut minima inventore et nam."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp2 := new(CreateClusterCommand)
	sub = &cobra.Command{
		Use:   `cluster ["/v1/projects/PROJECT/ns/NS/cluster"]`,
		Short: `Manage {create, delete}, and get cluster resources`,
		Long: `Manage {create, delete}, and get cluster resources

Payload example:

{
   "nodePoolSize": 10
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp3 := new(CreateNamespaceCommand)
	sub = &cobra.Command{
		Use:   `namespace ["/v1/projects/PROJECT/ns"]`,
		Short: `Manage {create, delete}, and get project's namespace(s)`,
		Long: `Manage {create, delete}, and get project's namespace(s)

Payload example:

{
   "name": "Non adipisci."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp4 := new(CreateProjectCommand)
	sub = &cobra.Command{
		Use:   `project ["/v1/projects"]`,
		Short: `Manage {create, delete} individual projects, read the list of all projects, read a specific project`,
		Long: `Manage {create, delete} individual projects, read the list of all projects, read a specific project

Payload example:

{
   "identity": "Perferendis enim."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `delete action`,
	}
	tmp5 := new(DeleteApplicationCommand)
	sub = &cobra.Command{
		Use:   `application ["/v1/projects/PROJECT/ns/NS/app/CHART"]`,
		Short: `Manage {create, delete}, and get namespaces's Application(s)`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp6 := new(DeleteClusterCommand)
	sub = &cobra.Command{
		Use:   `cluster ["/v1/projects/PROJECT/ns/NS/cluster"]`,
		Short: `Manage {create, delete}, and get cluster resources`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp7 := new(DeleteNamespaceCommand)
	sub = &cobra.Command{
		Use:   `namespace ["/v1/projects/PROJECT/ns/NS"]`,
		Short: `Manage {create, delete}, and get project's namespace(s)`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp8 := new(DeleteProjectCommand)
	sub = &cobra.Command{
		Use:   `project ["/v1/projects/PROJECT"]`,
		Short: `Manage {create, delete} individual projects, read the list of all projects, read a specific project`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "get",
		Short: `get action`,
	}
	tmp9 := new(GetApplicationCommand)
	sub = &cobra.Command{
		Use:   `application ["/v1/projects/PROJECT/ns/NS/app/CHART"]`,
		Short: `Manage {create, delete}, and get namespaces's Application(s)`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp9.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp10 := new(GetClusterCommand)
	sub = &cobra.Command{
		Use:   `cluster ["/v1/projects/PROJECT/ns/NS/cluster"]`,
		Short: `Manage {create, delete}, and get cluster resources`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp10.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp11 := new(GetNamespaceCommand)
	sub = &cobra.Command{
		Use:   `namespace ["/v1/projects/PROJECT/ns/NS"]`,
		Short: `Manage {create, delete}, and get project's namespace(s)`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp11.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp12 := new(GetProjectCommand)
	sub = &cobra.Command{
		Use:   `project ["/v1/projects/PROJECT"]`,
		Short: `Manage {create, delete} individual projects, read the list of all projects, read a specific project`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp12.Run(c, args) },
	}
	tmp12.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp12.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `list action`,
	}
	tmp13 := new(ListApplicationCommand)
	sub = &cobra.Command{
		Use:   `application ["/v1/projects/PROJECT/ns/NS/app"]`,
		Short: `Manage {create, delete}, and get namespaces's Application(s)`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp13.Run(c, args) },
	}
	tmp13.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp13.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp14 := new(ListNamespaceCommand)
	sub = &cobra.Command{
		Use:   `namespace ["/v1/projects/PROJECT/ns"]`,
		Short: `Manage {create, delete}, and get project's namespace(s)`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp14.Run(c, args) },
	}
	tmp14.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp14.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp15 := new(ListProjectCommand)
	sub = &cobra.Command{
		Use:   `project ["/v1/projects"]`,
		Short: `Manage {create, delete} individual projects, read the list of all projects, read a specific project`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp15.Run(c, args) },
	}
	tmp15.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp15.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)

	dl := new(DownloadCommand)
	dlc := &cobra.Command{
		Use:   "download [PATH]",
		Short: "Download file with given path",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dl.Run(c, args)
		},
	}
	dlc.Flags().StringVar(&dl.OutFile, "out", "", "Output file")
	app.AddCommand(dlc)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run downloads files with given paths.
func (cmd *DownloadCommand) Run(c *client.Client, args []string) error {
	var (
		fnf func(context.Context, string) (int64, error)
		fnd func(context.Context, string, string) (int64, error)

		rpath   = args[0]
		outfile = cmd.OutFile
		logger  = goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
		ctx     = goa.WithLogger(context.Background(), logger)
		err     error
	)

	if rpath[0] != '/' {
		rpath = "/" + rpath
	}
	if rpath == "/openapi" {
		fnf = c.DownloadOpenapi
		if outfile == "" {
			outfile = "swagger.json"
		}
		goto found
	}
	if rpath == "/openapi.json" {
		fnf = c.DownloadOpenapiJSON
		if outfile == "" {
			outfile = "swagger.json"
		}
		goto found
	}
	if rpath == "/openapi.yaml" {
		fnf = c.DownloadOpenapiYaml
		if outfile == "" {
			outfile = "swagger.yaml"
		}
		goto found
	}
	if rpath == "/swagger" {
		fnf = c.DownloadSwagger
		if outfile == "" {
			outfile = "swagger.json"
		}
		goto found
	}
	if rpath == "/swagger.json" {
		fnf = c.DownloadSwaggerJSON
		if outfile == "" {
			outfile = "swagger.json"
		}
		goto found
	}
	if rpath == "/swagger.yaml" {
		fnf = c.DownloadSwaggerYaml
		if outfile == "" {
			outfile = "swagger.yaml"
		}
		goto found
	}
	return fmt.Errorf("don't know how to download %s", rpath)
found:
	ctx = goa.WithLogContext(ctx, "file", outfile)
	if fnf != nil {
		_, err = fnf(ctx, outfile)
	} else {
		_, err = fnd(ctx, rpath, outfile)
	}
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	return nil
}

// Run makes the HTTP request corresponding to the CreateApplicationCommand command.
func (cmd *CreateApplicationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/projects/%v/ns/%v/app", url.QueryEscape(cmd.Project), url.QueryEscape(cmd.Ns))
	}
	var payload client.ApplicationPostBody
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateApplication(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateApplicationCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var ns string
	cc.Flags().StringVar(&cmd.Ns, "ns", ns, `namespace identifier`)
	var project string
	cc.Flags().StringVar(&cmd.Project, "project", project, `project name`)
}

// Run makes the HTTP request corresponding to the DeleteApplicationCommand command.
func (cmd *DeleteApplicationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/projects/%v/ns/%v/app/%v", url.QueryEscape(cmd.Project), url.QueryEscape(cmd.Ns), url.QueryEscape(cmd.Chart))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteApplication(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteApplicationCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var chart string
	cc.Flags().StringVar(&cmd.Chart, "chart", chart, ``)
	var ns string
	cc.Flags().StringVar(&cmd.Ns, "ns", ns, `namespace identifier`)
	var project string
	cc.Flags().StringVar(&cmd.Project, "project", project, `project name`)
}

// Run makes the HTTP request corresponding to the GetApplicationCommand command.
func (cmd *GetApplicationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/projects/%v/ns/%v/app/%v", url.QueryEscape(cmd.Project), url.QueryEscape(cmd.Ns), url.QueryEscape(cmd.Chart))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetApplication(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetApplicationCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var chart string
	cc.Flags().StringVar(&cmd.Chart, "chart", chart, ``)
	var ns string
	cc.Flags().StringVar(&cmd.Ns, "ns", ns, `namespace identifier`)
	var project string
	cc.Flags().StringVar(&cmd.Project, "project", project, `project name`)
}

// Run makes the HTTP request corresponding to the ListApplicationCommand command.
func (cmd *ListApplicationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/projects/%v/ns/%v/app", url.QueryEscape(cmd.Project), url.QueryEscape(cmd.Ns))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListApplication(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListApplicationCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var ns string
	cc.Flags().StringVar(&cmd.Ns, "ns", ns, `namespace identifier`)
	var project string
	cc.Flags().StringVar(&cmd.Project, "project", project, `project name`)
}

// Run makes the HTTP request corresponding to the CreateClusterCommand command.
func (cmd *CreateClusterCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/projects/%v/ns/%v/cluster", url.QueryEscape(cmd.Project), url.QueryEscape(cmd.Ns))
	}
	var payload client.CluterPostBody
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateCluster(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateClusterCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var ns string
	cc.Flags().StringVar(&cmd.Ns, "ns", ns, `namespace identifier`)
	var project string
	cc.Flags().StringVar(&cmd.Project, "project", project, `project name`)
}

// Run makes the HTTP request corresponding to the DeleteClusterCommand command.
func (cmd *DeleteClusterCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/projects/%v/ns/%v/cluster", url.QueryEscape(cmd.Project), url.QueryEscape(cmd.Ns))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteCluster(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteClusterCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var ns string
	cc.Flags().StringVar(&cmd.Ns, "ns", ns, `namespace identifier`)
	var project string
	cc.Flags().StringVar(&cmd.Project, "project", project, `project name`)
}

// Run makes the HTTP request corresponding to the GetClusterCommand command.
func (cmd *GetClusterCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/projects/%v/ns/%v/cluster", url.QueryEscape(cmd.Project), url.QueryEscape(cmd.Ns))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetCluster(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetClusterCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var ns string
	cc.Flags().StringVar(&cmd.Ns, "ns", ns, `namespace identifier`)
	var project string
	cc.Flags().StringVar(&cmd.Project, "project", project, `project name`)
}

// Run makes the HTTP request corresponding to the CreateNamespaceCommand command.
func (cmd *CreateNamespaceCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/projects/%v/ns", url.QueryEscape(cmd.Project))
	}
	var payload client.CreateNamespacePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateNamespace(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateNamespaceCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var project string
	cc.Flags().StringVar(&cmd.Project, "project", project, `project name`)
}

// Run makes the HTTP request corresponding to the DeleteNamespaceCommand command.
func (cmd *DeleteNamespaceCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/projects/%v/ns/%v", url.QueryEscape(cmd.Project), url.QueryEscape(cmd.Ns))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteNamespace(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteNamespaceCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var ns string
	cc.Flags().StringVar(&cmd.Ns, "ns", ns, `namespace identifier`)
	var project string
	cc.Flags().StringVar(&cmd.Project, "project", project, `project name`)
}

// Run makes the HTTP request corresponding to the GetNamespaceCommand command.
func (cmd *GetNamespaceCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/projects/%v/ns/%v", url.QueryEscape(cmd.Project), url.QueryEscape(cmd.Ns))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetNamespace(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetNamespaceCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var ns string
	cc.Flags().StringVar(&cmd.Ns, "ns", ns, `namespace identifier`)
	var project string
	cc.Flags().StringVar(&cmd.Project, "project", project, `project name`)
}

// Run makes the HTTP request corresponding to the ListNamespaceCommand command.
func (cmd *ListNamespaceCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/projects/%v/ns", url.QueryEscape(cmd.Project))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListNamespace(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListNamespaceCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var project string
	cc.Flags().StringVar(&cmd.Project, "project", project, `project name`)
}

// Run makes the HTTP request corresponding to the CreateProjectCommand command.
func (cmd *CreateProjectCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/projects"
	}
	var payload client.CreateProjectPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateProject(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateProjectCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the DeleteProjectCommand command.
func (cmd *DeleteProjectCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/projects/%v", url.QueryEscape(cmd.Project))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteProject(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteProjectCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var project string
	cc.Flags().StringVar(&cmd.Project, "project", project, `project name`)
}

// Run makes the HTTP request corresponding to the GetProjectCommand command.
func (cmd *GetProjectCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/projects/%v", url.QueryEscape(cmd.Project))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetProject(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetProjectCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var project string
	cc.Flags().StringVar(&cmd.Project, "project", project, `project name`)
}

// Run makes the HTTP request corresponding to the ListProjectCommand command.
func (cmd *ListProjectCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/projects"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListProject(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListProjectCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}
