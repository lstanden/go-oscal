/*
	This file was auto-generated with go-oscal.

	To regenerate:

	go-oscal \
		--input-file <path_to_oscal_json_schema_file> \
		--output-file <name_of_go_types_file> // the path to this file must already exist \
		--tags json,yaml // the tags to add to the Go structs \
		--pkg <name_of_your_go_package> // defaults to "main"

	For more information on how to use go-oscal: go-oscal --help

	Source: https://github.com/defenseunicorns/go-oscal
*/

package oscalTypes

type OscalModels struct {
	Catalog                   Catalog                   `json:"catalog" yaml:"catalog"`
	Profile                   Profile                   `json:"profile" yaml:"profile"`
	ComponentDefinition       ComponentDefinition       `json:"component-definition" yaml:"component-definition"`
	SystemSecurityPlan        SystemSecurityPlan        `json:"system-security-plan" yaml:"system-security-plan"`
	AssessmentPlan            AssessmentPlan            `json:"assessment-plan" yaml:"assessment-plan"`
	AssessmentResults         AssessmentResults         `json:"assessment-results" yaml:"assessment-results"`
	PlanOfActionAndMilestones PlanOfActionAndMilestones `json:"plan-of-action-and-milestones" yaml:"plan-of-action-and-milestones"`
}

type AssessmentPlan struct {
	AssessmentAssets   AssessmentAssets    `json:"assessment-assets,omitempty" yaml:"assessment-assets,omitempty"`
	AssessmentSubjects []AssessmentSubject `json:"assessment-subjects,omitempty" yaml:"assessment-subjects,omitempty"`
	BackMatter         BackMatter          `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	ImportSsp          ImportSsp           `json:"import-ssp" yaml:"import-ssp"`
	LocalDefinitions   LocalDefinitions    `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Metadata           Metadata            `json:"metadata" yaml:"metadata"`
	ReviewedControls   ReviewedControls    `json:"reviewed-controls" yaml:"reviewed-controls"`
	Tasks              []Task              `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	TermsAndConditions TermsAndConditions  `json:"terms-and-conditions,omitempty" yaml:"terms-and-conditions,omitempty"`
	UUID               string              `json:"uuid" yaml:"uuid"`
}

type AssessmentResults struct {
	BackMatter       BackMatter       `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	ImportAp         ImportAp         `json:"import-ap" yaml:"import-ap"`
	LocalDefinitions LocalDefinitions `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Metadata         Metadata         `json:"metadata" yaml:"metadata"`
	Results          []Result         `json:"results" yaml:"results"`
	UUID             string           `json:"uuid" yaml:"uuid"`
}

type Finding struct {
	Description                 string               `json:"description" yaml:"description"`
	ImplementationStatementUuid string               `json:"implementation-statement-uuid,omitempty" yaml:"implementation-statement-uuid,omitempty"`
	Links                       []Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Origins                     []Origin             `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props                       []Property           `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedObservations         []RelatedObservation `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	RelatedRisks                []AssociatedRisk     `json:"related-risks,omitempty" yaml:"related-risks,omitempty"`
	Remarks                     string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Target                      FindingTarget        `json:"target" yaml:"target"`
	Title                       string               `json:"title" yaml:"title"`
	UUID                        string               `json:"uuid" yaml:"uuid"`
}

type ImportAp struct {
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Result struct {
	AssessmentLog    AssessmentLog           `json:"assessment-log,omitempty" yaml:"assessment-log,omitempty"`
	Attestations     []AttestationStatements `json:"attestations,omitempty" yaml:"attestations,omitempty"`
	Description      string                  `json:"description" yaml:"description"`
	End              string                  `json:"end,omitempty" yaml:"end,omitempty"`
	Findings         []Finding               `json:"findings,omitempty" yaml:"findings,omitempty"`
	Links            []Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	LocalDefinitions LocalDefinitions        `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Observations     []Observation           `json:"observations,omitempty" yaml:"observations,omitempty"`
	Props            []Property              `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ReviewedControls ReviewedControls        `json:"reviewed-controls" yaml:"reviewed-controls"`
	Risks            []Risk                  `json:"risks,omitempty" yaml:"risks,omitempty"`
	Start            string                  `json:"start" yaml:"start"`
	Title            string                  `json:"title" yaml:"title"`
	UUID             string                  `json:"uuid" yaml:"uuid"`
}

type Activity struct {
	Description      string            `json:"description" yaml:"description"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedControls  ReviewedControls  `json:"related-controls,omitempty" yaml:"related-controls,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Steps            []Step            `json:"steps,omitempty" yaml:"steps,omitempty"`
	Title            string            `json:"title,omitempty" yaml:"title,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type AssessmentAssets struct {
	AssessmentPlatforms []AssessmentPlatform `json:"assessment-platforms" yaml:"assessment-platforms"`
	Components          []SystemComponent    `json:"components,omitempty" yaml:"components,omitempty"`
}

type AssessmentPart struct {
	Class string           `json:"class,omitempty" yaml:"class,omitempty"`
	Links []Link           `json:"links,omitempty" yaml:"links,omitempty"`
	Name  string           `json:"name" yaml:"name"`
	Ns    string           `json:"ns,omitempty" yaml:"ns,omitempty"`
	Parts []AssessmentPart `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props []Property       `json:"props,omitempty" yaml:"props,omitempty"`
	Prose string           `json:"prose,omitempty" yaml:"prose,omitempty"`
	Title string           `json:"title,omitempty" yaml:"title,omitempty"`
	UUID  string           `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type AssessmentSubject struct {
	Description     string              `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeSubjects []SelectSubjectById `json:"exclude-subjects,omitempty" yaml:"exclude-subjects,omitempty"`
	IncludeAll      IncludeAll          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeSubjects []SelectSubjectById `json:"include-subjects,omitempty" yaml:"include-subjects,omitempty"`
	Links           []Link              `json:"links,omitempty" yaml:"links,omitempty"`
	Props           []Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks         string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Type            string              `json:"type" yaml:"type"`
}

type Characterization struct {
	Facets []Facet    `json:"facets" yaml:"facets"`
	Links  []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Origin Origin     `json:"origin" yaml:"origin"`
	Props  []Property `json:"props,omitempty" yaml:"props,omitempty"`
}

type FindingTarget struct {
	Description          string               `json:"description,omitempty" yaml:"description,omitempty"`
	ImplementationStatus ImplementationStatus `json:"implementation-status,omitempty" yaml:"implementation-status,omitempty"`
	Links                []Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Props                []Property           `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Status               Status               `json:"status" yaml:"status"`
	TargetId             string               `json:"target-id" yaml:"target-id"`
	Title                string               `json:"title,omitempty" yaml:"title,omitempty"`
	Type                 string               `json:"type" yaml:"type"`
}

type ImportSsp struct {
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type LocalObjective struct {
	ControlId   string     `json:"control-id" yaml:"control-id"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Parts       []Part     `json:"parts" yaml:"parts"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type LoggedBy struct {
	PartyUuid string `json:"party-uuid" yaml:"party-uuid"`
	RoleId    string `json:"role-id,omitempty" yaml:"role-id,omitempty"`
}

type Observation struct {
	Collected        string             `json:"collected" yaml:"collected"`
	Description      string             `json:"description" yaml:"description"`
	Expires          string             `json:"expires,omitempty" yaml:"expires,omitempty"`
	Links            []Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Methods          []string           `json:"methods" yaml:"methods"`
	Origins          []Origin           `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props            []Property         `json:"props,omitempty" yaml:"props,omitempty"`
	RelevantEvidence []RelevantEvidence `json:"relevant-evidence,omitempty" yaml:"relevant-evidence,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Subjects         []SubjectReference `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	Title            string             `json:"title,omitempty" yaml:"title,omitempty"`
	Types            []string           `json:"types,omitempty" yaml:"types,omitempty"`
	UUID             string             `json:"uuid" yaml:"uuid"`
}

type Origin struct {
	Actors       []OriginActor `json:"actors" yaml:"actors"`
	RelatedTasks []RelatedTask `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
}

type OriginActor struct {
	ActorUuid string     `json:"actor-uuid" yaml:"actor-uuid"`
	Links     []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props     []Property `json:"props,omitempty" yaml:"props,omitempty"`
	RoleId    string     `json:"role-id,omitempty" yaml:"role-id,omitempty"`
	Type      string     `json:"type" yaml:"type"`
}

type RelatedTask struct {
	IdentifiedSubject  IdentifiedSubject   `json:"identified-subject,omitempty" yaml:"identified-subject,omitempty"`
	Links              []Link              `json:"links,omitempty" yaml:"links,omitempty"`
	Props              []Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties []ResponsibleParty  `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Subjects           []AssessmentSubject `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	TaskUuid           string              `json:"task-uuid" yaml:"task-uuid"`
}

type Response struct {
	Description    string          `json:"description" yaml:"description"`
	Lifecycle      string          `json:"lifecycle" yaml:"lifecycle"`
	Links          []Link          `json:"links,omitempty" yaml:"links,omitempty"`
	Origins        []Origin        `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props          []Property      `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks        string          `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RequiredAssets []RequiredAsset `json:"required-assets,omitempty" yaml:"required-assets,omitempty"`
	Tasks          []Task          `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Title          string          `json:"title" yaml:"title"`
	UUID           string          `json:"uuid" yaml:"uuid"`
}

type ReviewedControls struct {
	ControlObjectiveSelections []ReferencedControlObjectives `json:"control-objective-selections,omitempty" yaml:"control-objective-selections,omitempty"`
	ControlSelections          []AssessedControls            `json:"control-selections" yaml:"control-selections"`
	Description                string                        `json:"description,omitempty" yaml:"description,omitempty"`
	Links                      []Link                        `json:"links,omitempty" yaml:"links,omitempty"`
	Props                      []Property                    `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks                    string                        `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Risk struct {
	Characterizations   []Characterization   `json:"characterizations,omitempty" yaml:"characterizations,omitempty"`
	Deadline            string               `json:"deadline,omitempty" yaml:"deadline,omitempty"`
	Description         string               `json:"description" yaml:"description"`
	Links               []Link               `json:"links,omitempty" yaml:"links,omitempty"`
	MitigatingFactors   []MitigatingFactor   `json:"mitigating-factors,omitempty" yaml:"mitigating-factors,omitempty"`
	Origins             []Origin             `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props               []Property           `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedObservations []RelatedObservation `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	Remediations        []Response           `json:"remediations,omitempty" yaml:"remediations,omitempty"`
	RiskLog             RiskLog              `json:"risk-log,omitempty" yaml:"risk-log,omitempty"`
	Statement           string               `json:"statement" yaml:"statement"`
	Status              string               `json:"status" yaml:"status"`
	ThreatIds           []ThreatId           `json:"threat-ids,omitempty" yaml:"threat-ids,omitempty"`
	Title               string               `json:"title" yaml:"title"`
	UUID                string               `json:"uuid" yaml:"uuid"`
}

type SelectControlById struct {
	ControlId    string   `json:"control-id" yaml:"control-id"`
	StatementIds []string `json:"statement-ids,omitempty" yaml:"statement-ids,omitempty"`
}

type SelectObjectiveById struct {
	ObjectiveId string `json:"objective-id" yaml:"objective-id"`
}

type SelectSubjectById struct {
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	SubjectUuid string     `json:"subject-uuid" yaml:"subject-uuid"`
	Type        string     `json:"type" yaml:"type"`
}

type SubjectReference struct {
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	SubjectUuid string     `json:"subject-uuid" yaml:"subject-uuid"`
	Title       string     `json:"title,omitempty" yaml:"title,omitempty"`
	Type        string     `json:"type" yaml:"type"`
}

type Task struct {
	AssociatedActivities []AssociatedActivity `json:"associated-activities,omitempty" yaml:"associated-activities,omitempty"`
	Dependencies         []TaskDependency     `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	Description          string               `json:"description,omitempty" yaml:"description,omitempty"`
	Links                []Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Props                []Property           `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles     []ResponsibleRole    `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Subjects             []AssessmentSubject  `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	Tasks                []Task               `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Timing               Timing               `json:"timing,omitempty" yaml:"timing,omitempty"`
	Title                string               `json:"title" yaml:"title"`
	Type                 string               `json:"type" yaml:"type"`
	UUID                 string               `json:"uuid" yaml:"uuid"`
}

type IncludeAll struct {
}

type Parameter struct {
	Class       string                `json:"class,omitempty" yaml:"class,omitempty"`
	Constraints []ParameterConstraint `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	DependsOn   string                `json:"depends-on,omitempty" yaml:"depends-on,omitempty"`
	Guidelines  []ParameterGuideline  `json:"guidelines,omitempty" yaml:"guidelines,omitempty"`
	ID          string                `json:"id" yaml:"id"`
	Label       string                `json:"label,omitempty" yaml:"label,omitempty"`
	Links       []Link                `json:"links,omitempty" yaml:"links,omitempty"`
	Props       []Property            `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Select      ParameterSelection    `json:"select,omitempty" yaml:"select,omitempty"`
	Usage       string                `json:"usage,omitempty" yaml:"usage,omitempty"`
	Values      []string              `json:"values,omitempty" yaml:"values,omitempty"`
}

type ParameterConstraint struct {
	Description string           `json:"description,omitempty" yaml:"description,omitempty"`
	Tests       []ConstraintTest `json:"tests,omitempty" yaml:"tests,omitempty"`
}

type ParameterGuideline struct {
	Prose string `json:"prose" yaml:"prose"`
}

type ParameterSelection struct {
	Choice  []string `json:"choice,omitempty" yaml:"choice,omitempty"`
	HowMany string   `json:"how-many,omitempty" yaml:"how-many,omitempty"`
}

type Part struct {
	Class string     `json:"class,omitempty" yaml:"class,omitempty"`
	ID    string     `json:"id,omitempty" yaml:"id,omitempty"`
	Links []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Name  string     `json:"name" yaml:"name"`
	Ns    string     `json:"ns,omitempty" yaml:"ns,omitempty"`
	Parts []Part     `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Prose string     `json:"prose,omitempty" yaml:"prose,omitempty"`
	Title string     `json:"title,omitempty" yaml:"title,omitempty"`
}

type Catalog struct {
	BackMatter BackMatter  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Controls   []Control   `json:"controls,omitempty" yaml:"controls,omitempty"`
	Groups     []Group     `json:"groups,omitempty" yaml:"groups,omitempty"`
	Metadata   Metadata    `json:"metadata" yaml:"metadata"`
	Params     []Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	UUID       string      `json:"uuid" yaml:"uuid"`
}

type Control struct {
	Class    string      `json:"class,omitempty" yaml:"class,omitempty"`
	Controls []Control   `json:"controls,omitempty" yaml:"controls,omitempty"`
	ID       string      `json:"id" yaml:"id"`
	Links    []Link      `json:"links,omitempty" yaml:"links,omitempty"`
	Params   []Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	Parts    []Part      `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props    []Property  `json:"props,omitempty" yaml:"props,omitempty"`
	Title    string      `json:"title" yaml:"title"`
}

type Group struct {
	Class    string      `json:"class,omitempty" yaml:"class,omitempty"`
	Controls []Control   `json:"controls,omitempty" yaml:"controls,omitempty"`
	Groups   []Group     `json:"groups,omitempty" yaml:"groups,omitempty"`
	ID       string      `json:"id,omitempty" yaml:"id,omitempty"`
	Links    []Link      `json:"links,omitempty" yaml:"links,omitempty"`
	Params   []Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	Parts    []Part      `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props    []Property  `json:"props,omitempty" yaml:"props,omitempty"`
	Title    string      `json:"title" yaml:"title"`
}

type Capability struct {
	ControlImplementations []ControlImplementation `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Description            string                  `json:"description" yaml:"description"`
	IncorporatesComponents []IncorporatesComponent `json:"incorporates-components,omitempty" yaml:"incorporates-components,omitempty"`
	Links                  []Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	Name                   string                  `json:"name" yaml:"name"`
	Props                  []Property              `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks                string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID                   string                  `json:"uuid" yaml:"uuid"`
}

type ComponentDefinition struct {
	BackMatter                 BackMatter                  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Capabilities               []Capability                `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
	Components                 []DefinedComponent          `json:"components,omitempty" yaml:"components,omitempty"`
	ImportComponentDefinitions []ImportComponentDefinition `json:"import-component-definitions,omitempty" yaml:"import-component-definitions,omitempty"`
	Metadata                   Metadata                    `json:"metadata" yaml:"metadata"`
	UUID                       string                      `json:"uuid" yaml:"uuid"`
}

type ControlImplementation struct {
	Description             string                   `json:"description" yaml:"description"`
	ImplementedRequirements []ImplementedRequirement `json:"implemented-requirements" yaml:"implemented-requirements"`
	Links                   []Link                   `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   []Property               `json:"props,omitempty" yaml:"props,omitempty"`
	SetParameters           []SetParameter           `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Source                  string                   `json:"source" yaml:"source"`
	UUID                    string                   `json:"uuid" yaml:"uuid"`
}

type DefinedComponent struct {
	ControlImplementations []ControlImplementation `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Description            string                  `json:"description" yaml:"description"`
	Links                  []Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	Props                  []Property              `json:"props,omitempty" yaml:"props,omitempty"`
	Protocols              []Protocol              `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	Purpose                string                  `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Remarks                string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles       []ResponsibleRole       `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Title                  string                  `json:"title" yaml:"title"`
	Type                   string                  `json:"type" yaml:"type"`
	UUID                   string                  `json:"uuid" yaml:"uuid"`
}

type ImplementedRequirement struct {
	ControlId        string            `json:"control-id" yaml:"control-id"`
	Description      string            `json:"description" yaml:"description"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	SetParameters    []SetParameter    `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Statements       []Statement       `json:"statements,omitempty" yaml:"statements,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type ImportComponentDefinition struct {
	Href string `json:"href" yaml:"href"`
}

type IncorporatesComponent struct {
	ComponentUuid string `json:"component-uuid" yaml:"component-uuid"`
	Description   string `json:"description" yaml:"description"`
}

type Statement struct {
	Description      string            `json:"description" yaml:"description"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	StatementId      string            `json:"statement-id" yaml:"statement-id"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type AuthorizedPrivilege struct {
	Description        string   `json:"description,omitempty" yaml:"description,omitempty"`
	FunctionsPerformed []string `json:"functions-performed" yaml:"functions-performed"`
	Title              string   `json:"title" yaml:"title"`
}

type ImplementationStatus struct {
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string `json:"state" yaml:"state"`
}

type InventoryItem struct {
	Description           string                 `json:"description" yaml:"description"`
	ImplementedComponents []ImplementedComponent `json:"implemented-components,omitempty" yaml:"implemented-components,omitempty"`
	Links                 []Link                 `json:"links,omitempty" yaml:"links,omitempty"`
	Props                 []Property             `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks               string                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties    []ResponsibleParty     `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	UUID                  string                 `json:"uuid" yaml:"uuid"`
}

type PortRange struct {
	End       int    `json:"end,omitempty" yaml:"end,omitempty"`
	Start     int    `json:"start,omitempty" yaml:"start,omitempty"`
	Transport string `json:"transport,omitempty" yaml:"transport,omitempty"`
}

type Protocol struct {
	Name       string      `json:"name" yaml:"name"`
	PortRanges []PortRange `json:"port-ranges,omitempty" yaml:"port-ranges,omitempty"`
	Title      string      `json:"title,omitempty" yaml:"title,omitempty"`
	UUID       string      `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type SetParameter struct {
	ParamId string   `json:"param-id" yaml:"param-id"`
	Remarks string   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Values  []string `json:"values" yaml:"values"`
}

type SystemComponent struct {
	Description      string            `json:"description" yaml:"description"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Protocols        []Protocol        `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	Purpose          string            `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Status           Status            `json:"status" yaml:"status"`
	Title            string            `json:"title" yaml:"title"`
	Type             string            `json:"type" yaml:"type"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type SystemUser struct {
	AuthorizedPrivileges []AuthorizedPrivilege `json:"authorized-privileges,omitempty" yaml:"authorized-privileges,omitempty"`
	Description          string                `json:"description,omitempty" yaml:"description,omitempty"`
	Links                []Link                `json:"links,omitempty" yaml:"links,omitempty"`
	Props                []Property            `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleIds              []string              `json:"role-ids,omitempty" yaml:"role-ids,omitempty"`
	ShortName            string                `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	Title                string                `json:"title,omitempty" yaml:"title,omitempty"`
	UUID                 string                `json:"uuid" yaml:"uuid"`
}

type Address struct {
	AddrLines  []string `json:"addr-lines,omitempty" yaml:"addr-lines,omitempty"`
	City       string   `json:"city,omitempty" yaml:"city,omitempty"`
	Country    string   `json:"country,omitempty" yaml:"country,omitempty"`
	PostalCode string   `json:"postal-code,omitempty" yaml:"postal-code,omitempty"`
	State      string   `json:"state,omitempty" yaml:"state,omitempty"`
	Type       string   `json:"type,omitempty" yaml:"type,omitempty"`
}

type BackMatter struct {
	Resources []Resource `json:"resources,omitempty" yaml:"resources,omitempty"`
}

type Link struct {
	Href      string `json:"href" yaml:"href"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Rel       string `json:"rel,omitempty" yaml:"rel,omitempty"`
	Text      string `json:"text,omitempty" yaml:"text,omitempty"`
}

type Location struct {
	Address          Address           `json:"address" yaml:"address"`
	EmailAddresses   []string          `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	TelephoneNumbers []TelephoneNumber `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	Title            string            `json:"title,omitempty" yaml:"title,omitempty"`
	Urls             []string          `json:"urls,omitempty" yaml:"urls,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type Metadata struct {
	DocumentIds        []DocumentId       `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	LastModified       string             `json:"last-modified" yaml:"last-modified"`
	Links              []Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Locations          []Location         `json:"locations,omitempty" yaml:"locations,omitempty"`
	OscalVersion       string             `json:"oscal-version" yaml:"oscal-version"`
	Parties            []Party            `json:"parties,omitempty" yaml:"parties,omitempty"`
	Props              []Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Published          string             `json:"published,omitempty" yaml:"published,omitempty"`
	Remarks            string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties []ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Revisions          []Revision         `json:"revisions,omitempty" yaml:"revisions,omitempty"`
	Roles              []Role             `json:"roles,omitempty" yaml:"roles,omitempty"`
	Title              string             `json:"title" yaml:"title"`
	Version            string             `json:"version" yaml:"version"`
}

type Party struct {
	Addresses             []Address                 `json:"addresses,omitempty" yaml:"addresses,omitempty"`
	EmailAddresses        []string                  `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	ExternalIds           []PartyExternalIdentifier `json:"external-ids,omitempty" yaml:"external-ids,omitempty"`
	Links                 []Link                    `json:"links,omitempty" yaml:"links,omitempty"`
	LocationUuids         []string                  `json:"location-uuids,omitempty" yaml:"location-uuids,omitempty"`
	MemberOfOrganizations []string                  `json:"member-of-organizations,omitempty" yaml:"member-of-organizations,omitempty"`
	Name                  string                    `json:"name,omitempty" yaml:"name,omitempty"`
	Props                 []Property                `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks               string                    `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ShortName             string                    `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	TelephoneNumbers      []TelephoneNumber         `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	Type                  string                    `json:"type" yaml:"type"`
	UUID                  string                    `json:"uuid" yaml:"uuid"`
}

type Property struct {
	Class   string `json:"class,omitempty" yaml:"class,omitempty"`
	Name    string `json:"name" yaml:"name"`
	Ns      string `json:"ns,omitempty" yaml:"ns,omitempty"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID    string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Value   string `json:"value" yaml:"value"`
}

type ResponsibleParty struct {
	Links      []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuids []string   `json:"party-uuids" yaml:"party-uuids"`
	Props      []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks    string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleId     string     `json:"role-id" yaml:"role-id"`
}

type ResponsibleRole struct {
	Links      []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuids []string   `json:"party-uuids,omitempty" yaml:"party-uuids,omitempty"`
	Props      []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks    string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleId     string     `json:"role-id" yaml:"role-id"`
}

type Revision struct {
	LastModified string     `json:"last-modified,omitempty" yaml:"last-modified,omitempty"`
	Links        []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	OscalVersion string     `json:"oscal-version,omitempty" yaml:"oscal-version,omitempty"`
	Props        []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Published    string     `json:"published,omitempty" yaml:"published,omitempty"`
	Remarks      string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title        string     `json:"title,omitempty" yaml:"title,omitempty"`
	Version      string     `json:"version" yaml:"version"`
}

type Role struct {
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	ID          string     `json:"id" yaml:"id"`
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ShortName   string     `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	Title       string     `json:"title" yaml:"title"`
}

type LocalDefinitions struct {
	Components     []SystemComponent `json:"components,omitempty" yaml:"components,omitempty"`
	InventoryItems []InventoryItem   `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	Remarks        string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type PlanOfActionAndMilestones struct {
	BackMatter       BackMatter       `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	ImportSsp        ImportSsp        `json:"import-ssp,omitempty" yaml:"import-ssp,omitempty"`
	LocalDefinitions LocalDefinitions `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Metadata         Metadata         `json:"metadata" yaml:"metadata"`
	Observations     []Observation    `json:"observations,omitempty" yaml:"observations,omitempty"`
	PoamItems        []PoamItem       `json:"poam-items" yaml:"poam-items"`
	Risks            []Risk           `json:"risks,omitempty" yaml:"risks,omitempty"`
	SystemId         SystemId         `json:"system-id,omitempty" yaml:"system-id,omitempty"`
	UUID             string           `json:"uuid" yaml:"uuid"`
}

type PoamItem struct {
	Description         string               `json:"description" yaml:"description"`
	Links               []Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Origins             []Origin             `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props               []Property           `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedObservations []RelatedObservation `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	RelatedRisks        []AssociatedRisk     `json:"related-risks,omitempty" yaml:"related-risks,omitempty"`
	Remarks             string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title               string               `json:"title" yaml:"title"`
	UUID                string               `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type Add struct {
	ById     string      `json:"by-id,omitempty" yaml:"by-id,omitempty"`
	Links    []Link      `json:"links,omitempty" yaml:"links,omitempty"`
	Params   []Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	Parts    []Part      `json:"parts,omitempty" yaml:"parts,omitempty"`
	Position string      `json:"position,omitempty" yaml:"position,omitempty"`
	Props    []Property  `json:"props,omitempty" yaml:"props,omitempty"`
	Title    string      `json:"title,omitempty" yaml:"title,omitempty"`
}

type Alter struct {
	Adds      []Add    `json:"adds,omitempty" yaml:"adds,omitempty"`
	ControlId string   `json:"control-id" yaml:"control-id"`
	Removes   []Remove `json:"removes,omitempty" yaml:"removes,omitempty"`
}

type AssemblyOscalProfileGroup struct {
	Class          string           `json:"class,omitempty" yaml:"class,omitempty"`
	Groups         []Group          `json:"groups,omitempty" yaml:"groups,omitempty"`
	ID             string           `json:"id,omitempty" yaml:"id,omitempty"`
	InsertControls []InsertControls `json:"insert-controls,omitempty" yaml:"insert-controls,omitempty"`
	Links          []Link           `json:"links,omitempty" yaml:"links,omitempty"`
	Params         []Parameter      `json:"params,omitempty" yaml:"params,omitempty"`
	Parts          []Part           `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props          []Property       `json:"props,omitempty" yaml:"props,omitempty"`
	Title          string           `json:"title" yaml:"title"`
}

type Import struct {
	ExcludeControls []SelectControlById `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
	Href            string              `json:"href" yaml:"href"`
	IncludeAll      IncludeAll          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls []SelectControlById `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
}

type InsertControls struct {
	ExcludeControls []SelectControlById `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
	IncludeAll      IncludeAll          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls []SelectControlById `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
	Order           string              `json:"order,omitempty" yaml:"order,omitempty"`
}

type Merge struct {
	AsIs    bool    `json:"as-is,omitempty" yaml:"as-is,omitempty"`
	Combine Combine `json:"combine,omitempty" yaml:"combine,omitempty"`
	Custom  Custom  `json:"custom,omitempty" yaml:"custom,omitempty"`
	Flat    Flat    `json:"flat,omitempty" yaml:"flat,omitempty"`
}

type Modify struct {
	Alters        []Alter            `json:"alters,omitempty" yaml:"alters,omitempty"`
	SetParameters []ParameterSetting `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
}

type Profile struct {
	BackMatter BackMatter `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Imports    []Import   `json:"imports" yaml:"imports"`
	Merge      Merge      `json:"merge,omitempty" yaml:"merge,omitempty"`
	Metadata   Metadata   `json:"metadata" yaml:"metadata"`
	Modify     Modify     `json:"modify,omitempty" yaml:"modify,omitempty"`
	UUID       string     `json:"uuid" yaml:"uuid"`
}

type Remove struct {
	ByClass    string `json:"by-class,omitempty" yaml:"by-class,omitempty"`
	ById       string `json:"by-id,omitempty" yaml:"by-id,omitempty"`
	ByItemName string `json:"by-item-name,omitempty" yaml:"by-item-name,omitempty"`
	ByName     string `json:"by-name,omitempty" yaml:"by-name,omitempty"`
	ByNs       string `json:"by-ns,omitempty" yaml:"by-ns,omitempty"`
}

type AssemblyOscalProfileSelectControlById struct {
	Matching          []MatchControlsByPattern `json:"matching,omitempty" yaml:"matching,omitempty"`
	WithChildControls string                   `json:"with-child-controls,omitempty" yaml:"with-child-controls,omitempty"`
	WithIds           []string                 `json:"with-ids,omitempty" yaml:"with-ids,omitempty"`
}

type AuthorizationBoundary struct {
	Description string     `json:"description" yaml:"description"`
	Diagrams    []Diagram  `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ByComponent struct {
	ComponentUuid        string                                         `json:"component-uuid" yaml:"component-uuid"`
	Description          string                                         `json:"description" yaml:"description"`
	Export               Export                                         `json:"export,omitempty" yaml:"export,omitempty"`
	ImplementationStatus ImplementationStatus                           `json:"implementation-status,omitempty" yaml:"implementation-status,omitempty"`
	Inherited            []InheritedControlImplementation               `json:"inherited,omitempty" yaml:"inherited,omitempty"`
	Links                []Link                                         `json:"links,omitempty" yaml:"links,omitempty"`
	Props                []Property                                     `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              string                                         `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles     []ResponsibleRole                              `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Satisfied            []SatisfiedControlImplementationResponsibility `json:"satisfied,omitempty" yaml:"satisfied,omitempty"`
	SetParameters        []SetParameter                                 `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	UUID                 string                                         `json:"uuid" yaml:"uuid"`
}

type AssemblyOscalSspControlImplementation struct {
	Description             string                   `json:"description" yaml:"description"`
	ImplementedRequirements []ImplementedRequirement `json:"implemented-requirements" yaml:"implemented-requirements"`
	SetParameters           []SetParameter           `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
}

type DataFlow struct {
	Description string     `json:"description" yaml:"description"`
	Diagrams    []Diagram  `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Diagram struct {
	Caption     string     `json:"caption,omitempty" yaml:"caption,omitempty"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID        string     `json:"uuid" yaml:"uuid"`
}

type AssemblyOscalSspImplementedRequirement struct {
	ByComponents     []ByComponent     `json:"by-components,omitempty" yaml:"by-components,omitempty"`
	ControlId        string            `json:"control-id" yaml:"control-id"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	SetParameters    []SetParameter    `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Statements       []Statement       `json:"statements,omitempty" yaml:"statements,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type ImportProfile struct {
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type NetworkArchitecture struct {
	Description string     `json:"description" yaml:"description"`
	Diagrams    []Diagram  `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SecurityImpactLevel struct {
	SecurityObjectiveAvailability    string `json:"security-objective-availability" yaml:"security-objective-availability"`
	SecurityObjectiveConfidentiality string `json:"security-objective-confidentiality" yaml:"security-objective-confidentiality"`
	SecurityObjectiveIntegrity       string `json:"security-objective-integrity" yaml:"security-objective-integrity"`
}

type AssemblyOscalSspStatement struct {
	ByComponents     []ByComponent     `json:"by-components,omitempty" yaml:"by-components,omitempty"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	StatementId      string            `json:"statement-id" yaml:"statement-id"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type Status struct {
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string `json:"state" yaml:"state"`
}

type SystemCharacteristics struct {
	AuthorizationBoundary    AuthorizationBoundary `json:"authorization-boundary" yaml:"authorization-boundary"`
	DataFlow                 DataFlow              `json:"data-flow,omitempty" yaml:"data-flow,omitempty"`
	DateAuthorized           string                `json:"date-authorized,omitempty" yaml:"date-authorized,omitempty"`
	Description              string                `json:"description" yaml:"description"`
	Links                    []Link                `json:"links,omitempty" yaml:"links,omitempty"`
	NetworkArchitecture      NetworkArchitecture   `json:"network-architecture,omitempty" yaml:"network-architecture,omitempty"`
	Props                    []Property            `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks                  string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties       []ResponsibleParty    `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	SecurityImpactLevel      SecurityImpactLevel   `json:"security-impact-level" yaml:"security-impact-level"`
	SecuritySensitivityLevel string                `json:"security-sensitivity-level" yaml:"security-sensitivity-level"`
	Status                   Status                `json:"status" yaml:"status"`
	SystemIds                []SystemId            `json:"system-ids" yaml:"system-ids"`
	SystemInformation        SystemInformation     `json:"system-information" yaml:"system-information"`
	SystemName               string                `json:"system-name" yaml:"system-name"`
	SystemNameShort          string                `json:"system-name-short,omitempty" yaml:"system-name-short,omitempty"`
}

type SystemImplementation struct {
	Components              []SystemComponent        `json:"components" yaml:"components"`
	InventoryItems          []InventoryItem          `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	LeveragedAuthorizations []LeveragedAuthorization `json:"leveraged-authorizations,omitempty" yaml:"leveraged-authorizations,omitempty"`
	Links                   []Link                   `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   []Property               `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks                 string                   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Users                   []SystemUser             `json:"users" yaml:"users"`
}

type SystemInformation struct {
	InformationTypes []InformationType `json:"information-types" yaml:"information-types"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
}

type SystemSecurityPlan struct {
	BackMatter            BackMatter            `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	ControlImplementation ControlImplementation `json:"control-implementation" yaml:"control-implementation"`
	ImportProfile         ImportProfile         `json:"import-profile" yaml:"import-profile"`
	Metadata              Metadata              `json:"metadata" yaml:"metadata"`
	SystemCharacteristics SystemCharacteristics `json:"system-characteristics" yaml:"system-characteristics"`
	SystemImplementation  SystemImplementation  `json:"system-implementation" yaml:"system-implementation"`
	UUID                  string                `json:"uuid" yaml:"uuid"`
}

type ThreatId struct {
	Href   string `json:"href,omitempty" yaml:"href,omitempty"`
	ID     string `json:"id" yaml:"id"`
	System string `json:"system" yaml:"system"`
}

type SystemId struct {
	ID             string `json:"id" yaml:"id"`
	IdentifierType string `json:"identifier-type,omitempty" yaml:"identifier-type,omitempty"`
}

type DocumentId struct {
	Identifier string `json:"identifier" yaml:"identifier"`
	Scheme     string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
}

type Hash struct {
	Algorithm string `json:"algorithm" yaml:"algorithm"`
	Value     string `json:"value" yaml:"value"`
}

type TelephoneNumber struct {
	Number string `json:"number" yaml:"number"`
	Type   string `json:"type,omitempty" yaml:"type,omitempty"`
}

type AssessmentLog struct {
	Entries []AssessmentLogEntry `json:"entries" yaml:"entries"`
}

type AssessmentPlatform struct {
	Links          []Link          `json:"links,omitempty" yaml:"links,omitempty"`
	Props          []Property      `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks        string          `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title          string          `json:"title,omitempty" yaml:"title,omitempty"`
	UsesComponents []UsesComponent `json:"uses-components,omitempty" yaml:"uses-components,omitempty"`
	UUID           string          `json:"uuid" yaml:"uuid"`
}

type AssociatedActivity struct {
	ActivityUuid     string              `json:"activity-uuid" yaml:"activity-uuid"`
	Links            []Link              `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole   `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Subjects         []AssessmentSubject `json:"subjects" yaml:"subjects"`
}

type AtFrequency struct {
	Period int    `json:"period" yaml:"period"`
	Unit   string `json:"unit" yaml:"unit"`
}

type AttestationStatements struct {
	Parts              []AssessmentPart   `json:"parts" yaml:"parts"`
	ResponsibleParties []ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
}

type AvailabilityImpact struct {
	AdjustmentJustification string     `json:"adjustment-justification,omitempty" yaml:"adjustment-justification,omitempty"`
	Base                    string     `json:"base" yaml:"base"`
	Links                   []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Selected                string     `json:"selected,omitempty" yaml:"selected,omitempty"`
}

type Base64 struct {
	Filename  string `json:"filename,omitempty" yaml:"filename,omitempty"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Value     string `json:"value" yaml:"value"`
}

type InformationTypeCategorization struct {
	InformationTypeIds []string `json:"information-type-ids,omitempty" yaml:"information-type-ids,omitempty"`
	System             string   `json:"system" yaml:"system"`
}

type Citation struct {
	Links []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Text  string     `json:"text" yaml:"text"`
}

type Combine struct {
	Method string `json:"method,omitempty" yaml:"method,omitempty"`
}

type ConfidentialityImpact struct {
	AdjustmentJustification string     `json:"adjustment-justification,omitempty" yaml:"adjustment-justification,omitempty"`
	Base                    string     `json:"base" yaml:"base"`
	Links                   []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Selected                string     `json:"selected,omitempty" yaml:"selected,omitempty"`
}

type ReferencedControlObjectives struct {
	Description       string                `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeObjectives []SelectObjectiveById `json:"exclude-objectives,omitempty" yaml:"exclude-objectives,omitempty"`
	IncludeAll        IncludeAll            `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeObjectives []SelectObjectiveById `json:"include-objectives,omitempty" yaml:"include-objectives,omitempty"`
	Links             []Link                `json:"links,omitempty" yaml:"links,omitempty"`
	Props             []Property            `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks           string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type AssessedControls struct {
	Description     string              `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeControls []SelectControlById `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
	IncludeAll      IncludeAll          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls []SelectControlById `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
	Links           []Link              `json:"links,omitempty" yaml:"links,omitempty"`
	Props           []Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks         string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Custom struct {
	Groups         []Group          `json:"groups,omitempty" yaml:"groups,omitempty"`
	InsertControls []InsertControls `json:"insert-controls,omitempty" yaml:"insert-controls,omitempty"`
}

type TaskDependency struct {
	Remarks  string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	TaskUuid string `json:"task-uuid" yaml:"task-uuid"`
}

type AssessmentLogEntry struct {
	Description  string        `json:"description,omitempty" yaml:"description,omitempty"`
	End          string        `json:"end,omitempty" yaml:"end,omitempty"`
	Links        []Link        `json:"links,omitempty" yaml:"links,omitempty"`
	LoggedBy     []LoggedBy    `json:"logged-by,omitempty" yaml:"logged-by,omitempty"`
	Props        []Property    `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedTasks []RelatedTask `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
	Remarks      string        `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Start        string        `json:"start" yaml:"start"`
	Title        string        `json:"title,omitempty" yaml:"title,omitempty"`
	UUID         string        `json:"uuid" yaml:"uuid"`
}

type Export struct {
	Description      string                                `json:"description,omitempty" yaml:"description,omitempty"`
	Links            []Link                                `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property                            `json:"props,omitempty" yaml:"props,omitempty"`
	Provided         []ProvidedControlImplementation       `json:"provided,omitempty" yaml:"provided,omitempty"`
	Remarks          string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Responsibilities []ControlImplementationResponsibility `json:"responsibilities,omitempty" yaml:"responsibilities,omitempty"`
}

type PartyExternalIdentifier struct {
	ID     string `json:"id" yaml:"id"`
	Scheme string `json:"scheme" yaml:"scheme"`
}

type Facet struct {
	Links   []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Name    string     `json:"name" yaml:"name"`
	Props   []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	System  string     `json:"system" yaml:"system"`
	Value   string     `json:"value" yaml:"value"`
}

type Flat struct {
}

type IdentifiedSubject struct {
	SubjectPlaceholderUuid string              `json:"subject-placeholder-uuid" yaml:"subject-placeholder-uuid"`
	Subjects               []AssessmentSubject `json:"subjects" yaml:"subjects"`
}

type ImplementedComponent struct {
	ComponentUuid      string             `json:"component-uuid" yaml:"component-uuid"`
	Links              []Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              []Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties []ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
}

type InformationType struct {
	AvailabilityImpact    AvailabilityImpact              `json:"availability-impact" yaml:"availability-impact"`
	Categorizations       []InformationTypeCategorization `json:"categorizations,omitempty" yaml:"categorizations,omitempty"`
	ConfidentialityImpact ConfidentialityImpact           `json:"confidentiality-impact" yaml:"confidentiality-impact"`
	Description           string                          `json:"description" yaml:"description"`
	IntegrityImpact       IntegrityImpact                 `json:"integrity-impact" yaml:"integrity-impact"`
	Links                 []Link                          `json:"links,omitempty" yaml:"links,omitempty"`
	Props                 []Property                      `json:"props,omitempty" yaml:"props,omitempty"`
	Title                 string                          `json:"title" yaml:"title"`
	UUID                  string                          `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type InheritedControlImplementation struct {
	Description      string            `json:"description" yaml:"description"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	ProvidedUuid     string            `json:"provided-uuid,omitempty" yaml:"provided-uuid,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type IntegrityImpact struct {
	AdjustmentJustification string     `json:"adjustment-justification,omitempty" yaml:"adjustment-justification,omitempty"`
	Base                    string     `json:"base" yaml:"base"`
	Links                   []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Selected                string     `json:"selected,omitempty" yaml:"selected,omitempty"`
}

type LeveragedAuthorization struct {
	DateAuthorized string     `json:"date-authorized" yaml:"date-authorized"`
	Links          []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuid      string     `json:"party-uuid" yaml:"party-uuid"`
	Props          []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks        string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title          string     `json:"title" yaml:"title"`
	UUID           string     `json:"uuid" yaml:"uuid"`
}

type LocalDefinitions1 struct {
	Activities           []Activity        `json:"activities,omitempty" yaml:"activities,omitempty"`
	Components           []SystemComponent `json:"components,omitempty" yaml:"components,omitempty"`
	InventoryItems       []InventoryItem   `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	ObjectivesAndMethods []LocalObjective  `json:"objectives-and-methods,omitempty" yaml:"objectives-and-methods,omitempty"`
	Remarks              string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Users                []SystemUser      `json:"users,omitempty" yaml:"users,omitempty"`
}

type MatchControlsByPattern struct {
	Pattern string `json:"pattern,omitempty" yaml:"pattern,omitempty"`
}

type MitigatingFactor struct {
	Description        string             `json:"description" yaml:"description"`
	ImplementationUuid string             `json:"implementation-uuid,omitempty" yaml:"implementation-uuid,omitempty"`
	Links              []Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              []Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Subjects           []SubjectReference `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	UUID               string             `json:"uuid" yaml:"uuid"`
}

type OnDate struct {
	Date string `json:"date" yaml:"date"`
}

type Origins1 struct {
	Actors []OriginActor `json:"actors" yaml:"actors"`
}

type ProvidedControlImplementation struct {
	Description      string            `json:"description" yaml:"description"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type RelatedObservation struct {
	ObservationUuid string `json:"observation-uuid" yaml:"observation-uuid"`
}

type AssociatedRisk struct {
	RiskUuid string `json:"risk-uuid" yaml:"risk-uuid"`
}

type RelevantEvidence struct {
	Description string     `json:"description" yaml:"description"`
	Href        string     `json:"href,omitempty" yaml:"href,omitempty"`
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type RequiredAsset struct {
	Description string             `json:"description" yaml:"description"`
	Links       []Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props       []Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Subjects    []SubjectReference `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	Title       string             `json:"title,omitempty" yaml:"title,omitempty"`
	UUID        string             `json:"uuid" yaml:"uuid"`
}

type Resource struct {
	Base64      Base64         `json:"base64,omitempty" yaml:"base64,omitempty"`
	Citation    Citation       `json:"citation,omitempty" yaml:"citation,omitempty"`
	Description string         `json:"description,omitempty" yaml:"description,omitempty"`
	DocumentIds []DocumentId   `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	Props       []Property     `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string         `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Rlinks      []ResourceLink `json:"rlinks,omitempty" yaml:"rlinks,omitempty"`
	Title       string         `json:"title,omitempty" yaml:"title,omitempty"`
	UUID        string         `json:"uuid" yaml:"uuid"`
}

type ControlImplementationResponsibility struct {
	Description      string            `json:"description" yaml:"description"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	ProvidedUuid     string            `json:"provided-uuid,omitempty" yaml:"provided-uuid,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type RiskLog struct {
	Entries []AssessmentLogEntry `json:"entries" yaml:"entries"`
}

type ResourceLink struct {
	Hashes    []Hash `json:"hashes,omitempty" yaml:"hashes,omitempty"`
	Href      string `json:"href" yaml:"href"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
}

type SatisfiedControlImplementationResponsibility struct {
	Description        string            `json:"description" yaml:"description"`
	Links              []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props              []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibilityUuid string            `json:"responsibility-uuid,omitempty" yaml:"responsibility-uuid,omitempty"`
	ResponsibleRoles   []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	UUID               string            `json:"uuid" yaml:"uuid"`
}

type ParameterSetting struct {
	Class       string                `json:"class,omitempty" yaml:"class,omitempty"`
	Constraints []ParameterConstraint `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	DependsOn   string                `json:"depends-on,omitempty" yaml:"depends-on,omitempty"`
	Guidelines  []ParameterGuideline  `json:"guidelines,omitempty" yaml:"guidelines,omitempty"`
	Label       string                `json:"label,omitempty" yaml:"label,omitempty"`
	Links       []Link                `json:"links,omitempty" yaml:"links,omitempty"`
	ParamId     string                `json:"param-id" yaml:"param-id"`
	Props       []Property            `json:"props,omitempty" yaml:"props,omitempty"`
	Select      ParameterSelection    `json:"select,omitempty" yaml:"select,omitempty"`
	Usage       string                `json:"usage,omitempty" yaml:"usage,omitempty"`
	Values      []string              `json:"values,omitempty" yaml:"values,omitempty"`
}

type Status1 struct {
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string `json:"state" yaml:"state"`
}

type Step struct {
	Description      string            `json:"description" yaml:"description"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	ReviewedControls ReviewedControls  `json:"reviewed-controls,omitempty" yaml:"reviewed-controls,omitempty"`
	Title            string            `json:"title,omitempty" yaml:"title,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type TermsAndConditions struct {
	Parts []AssessmentPart `json:"parts,omitempty" yaml:"parts,omitempty"`
}

type ConstraintTest struct {
	Expression string `json:"expression" yaml:"expression"`
	Remarks    string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Timing struct {
	AtFrequency     AtFrequency     `json:"at-frequency,omitempty" yaml:"at-frequency,omitempty"`
	OnDate          OnDate          `json:"on-date,omitempty" yaml:"on-date,omitempty"`
	WithinDateRange WithinDateRange `json:"within-date-range,omitempty" yaml:"within-date-range,omitempty"`
}

type UsesComponent struct {
	ComponentUuid      string             `json:"component-uuid" yaml:"component-uuid"`
	Links              []Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              []Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties []ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
}

type WithinDateRange struct {
	End   string `json:"end" yaml:"end"`
	Start string `json:"start" yaml:"start"`
}
