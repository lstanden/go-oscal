/*
This file was auto-generated with go-oscal.

To regenerate:

	go-oscal generate \
		--input-file <path_to_oscal_json_schema_file> \
		--output-file <name_of_go_types_file> // the path to this file must already exist \
		--tags json,yaml // the tags to add to the Go structs \
		--pkg <name_of_your_go_package> // defaults to "main"

For more information on how to use go-oscal: go-oscal --help

Source: https://github.com/defenseunicorns/go-oscal
*/
package oscalTypes_1_0_4

import (
	"github.com/defenseunicorns/go-oscal/src/pkg/optional"
	"time"
)

type OscalModels = OscalCompleteSchema
type OscalCompleteSchema struct {
	AssessmentPlan            optional.Optional[AssessmentPlan]            `json:"assessment-plan,omitempty" yaml:"assessment-plan,omitempty"`
	AssessmentResults         optional.Optional[AssessmentResults]         `json:"assessment-results,omitempty" yaml:"assessment-results,omitempty"`
	Catalog                   optional.Optional[Catalog]                   `json:"catalog,omitempty" yaml:"catalog,omitempty"`
	ComponentDefinition       optional.Optional[ComponentDefinition]       `json:"component-definition,omitempty" yaml:"component-definition,omitempty"`
	PlanOfActionAndMilestones optional.Optional[PlanOfActionAndMilestones] `json:"plan-of-action-and-milestones,omitempty" yaml:"plan-of-action-and-milestones,omitempty"`
	Profile                   optional.Optional[Profile]                   `json:"profile,omitempty" yaml:"profile,omitempty"`
	SystemSecurityPlan        optional.Optional[SystemSecurityPlan]        `json:"system-security-plan,omitempty" yaml:"system-security-plan,omitempty"`
}

type AssessmentPlan struct {
	AssessmentAssets   optional.Optional[AssessmentAssets]                 `json:"assessment-assets,omitempty" yaml:"assessment-assets,omitempty"`
	AssessmentSubjects optional.Optional[[]AssessmentSubject]              `json:"assessment-subjects,omitempty" yaml:"assessment-subjects,omitempty"`
	BackMatter         optional.Optional[BackMatter]                       `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	ImportSsp          ImportSsp                                           `json:"import-ssp" yaml:"import-ssp"`
	LocalDefinitions   optional.Optional[LocalDefinitions]                 `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Metadata           Metadata                                            `json:"metadata" yaml:"metadata"`
	ReviewedControls   ReviewedControls                                    `json:"reviewed-controls" yaml:"reviewed-controls"`
	Tasks              optional.Optional[[]Task]                           `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	TermsAndConditions optional.Optional[AssessmentPlanTermsAndConditions] `json:"terms-and-conditions,omitempty" yaml:"terms-and-conditions,omitempty"`
	UUID               string                                              `json:"uuid" yaml:"uuid"`
}

type AssessmentResults struct {
	BackMatter       optional.Optional[BackMatter]       `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	ImportAp         ImportAp                            `json:"import-ap" yaml:"import-ap"`
	LocalDefinitions optional.Optional[LocalDefinitions] `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Metadata         Metadata                            `json:"metadata" yaml:"metadata"`
	Results          []Result                            `json:"results" yaml:"results"`
	UUID             string                              `json:"uuid" yaml:"uuid"`
}

type Catalog struct {
	BackMatter optional.Optional[BackMatter]  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Controls   optional.Optional[[]Control]   `json:"controls,omitempty" yaml:"controls,omitempty"`
	Groups     optional.Optional[[]Group]     `json:"groups,omitempty" yaml:"groups,omitempty"`
	Metadata   Metadata                       `json:"metadata" yaml:"metadata"`
	Params     optional.Optional[[]Parameter] `json:"params,omitempty" yaml:"params,omitempty"`
	UUID       string                         `json:"uuid" yaml:"uuid"`
}

type ComponentDefinition struct {
	BackMatter                 optional.Optional[BackMatter]                  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Capabilities               optional.Optional[[]Capability]                `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
	Components                 optional.Optional[[]DefinedComponent]          `json:"components,omitempty" yaml:"components,omitempty"`
	ImportComponentDefinitions optional.Optional[[]ImportComponentDefinition] `json:"import-component-definitions,omitempty" yaml:"import-component-definitions,omitempty"`
	Metadata                   Metadata                                       `json:"metadata" yaml:"metadata"`
	UUID                       string                                         `json:"uuid" yaml:"uuid"`
}

type PlanOfActionAndMilestones struct {
	BackMatter       optional.Optional[BackMatter]                                `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	ImportSsp        optional.Optional[ImportSsp]                                 `json:"import-ssp,omitempty" yaml:"import-ssp,omitempty"`
	LocalDefinitions optional.Optional[PlanOfActionAndMilestonesLocalDefinitions] `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Metadata         Metadata                                                     `json:"metadata" yaml:"metadata"`
	Observations     optional.Optional[[]Observation]                             `json:"observations,omitempty" yaml:"observations,omitempty"`
	PoamItems        []PoamItem                                                   `json:"poam-items" yaml:"poam-items"`
	Risks            optional.Optional[[]Risk]                                    `json:"risks,omitempty" yaml:"risks,omitempty"`
	SystemId         optional.Optional[SystemId]                                  `json:"system-id,omitempty" yaml:"system-id,omitempty"`
	UUID             string                                                       `json:"uuid" yaml:"uuid"`
}

type Profile struct {
	BackMatter optional.Optional[BackMatter] `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Imports    []Import                      `json:"imports" yaml:"imports"`
	Merge      optional.Optional[Merge]      `json:"merge,omitempty" yaml:"merge,omitempty"`
	Metadata   Metadata                      `json:"metadata" yaml:"metadata"`
	Modify     optional.Optional[Modify]     `json:"modify,omitempty" yaml:"modify,omitempty"`
	UUID       string                        `json:"uuid" yaml:"uuid"`
}

type SystemSecurityPlan struct {
	BackMatter            optional.Optional[BackMatter] `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	ControlImplementation ControlImplementation         `json:"control-implementation" yaml:"control-implementation"`
	ImportProfile         ImportProfile                 `json:"import-profile" yaml:"import-profile"`
	Metadata              Metadata                      `json:"metadata" yaml:"metadata"`
	SystemCharacteristics SystemCharacteristics         `json:"system-characteristics" yaml:"system-characteristics"`
	SystemImplementation  SystemImplementation          `json:"system-implementation" yaml:"system-implementation"`
	UUID                  string                        `json:"uuid" yaml:"uuid"`
}

type AssessmentAssets struct {
	AssessmentPlatforms []AssessmentPlatform                 `json:"assessment-platforms" yaml:"assessment-platforms"`
	Components          optional.Optional[[]SystemComponent] `json:"components,omitempty" yaml:"components,omitempty"`
}

type AssessmentSubject struct {
	Description     optional.Optional[string]              `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeSubjects optional.Optional[[]SelectSubjectById] `json:"exclude-subjects,omitempty" yaml:"exclude-subjects,omitempty"`
	IncludeAll      optional.Optional[IncludeAll]          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeSubjects optional.Optional[[]SelectSubjectById] `json:"include-subjects,omitempty" yaml:"include-subjects,omitempty"`
	Links           optional.Optional[[]Link]              `json:"links,omitempty" yaml:"links,omitempty"`
	Props           optional.Optional[[]Property]          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks         optional.Optional[string]              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Type            string                                 `json:"type" yaml:"type"`
}

type BackMatter struct {
	Resources optional.Optional[[]Resource] `json:"resources,omitempty" yaml:"resources,omitempty"`
}

type ImportSsp struct {
	Href    string                    `json:"href" yaml:"href"`
	Remarks optional.Optional[string] `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type LocalDefinitions struct {
	Activities           optional.Optional[[]Activity]        `json:"activities,omitempty" yaml:"activities,omitempty"`
	Components           optional.Optional[[]SystemComponent] `json:"components,omitempty" yaml:"components,omitempty"`
	InventoryItems       optional.Optional[[]InventoryItem]   `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	ObjectivesAndMethods optional.Optional[[]LocalObjective]  `json:"objectives-and-methods,omitempty" yaml:"objectives-and-methods,omitempty"`
	Remarks              optional.Optional[string]            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Users                optional.Optional[[]SystemUser]      `json:"users,omitempty" yaml:"users,omitempty"`
}

type Metadata struct {
	DocumentIds        optional.Optional[[]DocumentId]       `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	LastModified       time.Time                             `json:"last-modified" yaml:"last-modified"`
	Links              optional.Optional[[]Link]             `json:"links,omitempty" yaml:"links,omitempty"`
	Locations          optional.Optional[[]Location]         `json:"locations,omitempty" yaml:"locations,omitempty"`
	OscalVersion       string                                `json:"oscal-version" yaml:"oscal-version"`
	Parties            optional.Optional[[]Party]            `json:"parties,omitempty" yaml:"parties,omitempty"`
	Props              optional.Optional[[]Property]         `json:"props,omitempty" yaml:"props,omitempty"`
	Published          optional.Optional[time.Time]          `json:"published,omitempty" yaml:"published,omitempty"`
	Remarks            optional.Optional[string]             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties optional.Optional[[]ResponsibleParty] `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Revisions          optional.Optional[[]Revision]         `json:"revisions,omitempty" yaml:"revisions,omitempty"`
	Roles              optional.Optional[[]Role]             `json:"roles,omitempty" yaml:"roles,omitempty"`
	Title              string                                `json:"title" yaml:"title"`
	Version            string                                `json:"version" yaml:"version"`
}

type ReviewedControls struct {
	ControlObjectiveSelections optional.Optional[[]ReferencedControlObjectives] `json:"control-objective-selections,omitempty" yaml:"control-objective-selections,omitempty"`
	ControlSelections          []AssessedControls                               `json:"control-selections" yaml:"control-selections"`
	Description                optional.Optional[string]                        `json:"description,omitempty" yaml:"description,omitempty"`
	Links                      optional.Optional[[]Link]                        `json:"links,omitempty" yaml:"links,omitempty"`
	Props                      optional.Optional[[]Property]                    `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks                    optional.Optional[string]                        `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Task struct {
	AssociatedActivities optional.Optional[[]AssociatedActivity] `json:"associated-activities,omitempty" yaml:"associated-activities,omitempty"`
	Dependencies         optional.Optional[[]TaskDependency]     `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	Description          optional.Optional[string]               `json:"description,omitempty" yaml:"description,omitempty"`
	Links                optional.Optional[[]Link]               `json:"links,omitempty" yaml:"links,omitempty"`
	Props                optional.Optional[[]Property]           `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              optional.Optional[string]               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles     optional.Optional[[]ResponsibleRole]    `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Subjects             optional.Optional[[]AssessmentSubject]  `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	Tasks                optional.Optional[[]Task]               `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Timing               optional.Optional[EventTiming]          `json:"timing,omitempty" yaml:"timing,omitempty"`
	Title                string                                  `json:"title" yaml:"title"`
	Type                 string                                  `json:"type" yaml:"type"`
	UUID                 string                                  `json:"uuid" yaml:"uuid"`
}

type AssessmentPlanTermsAndConditions struct {
	Parts optional.Optional[[]AssessmentPart] `json:"parts,omitempty" yaml:"parts,omitempty"`
}

type ImportAp struct {
	Href    string                    `json:"href" yaml:"href"`
	Remarks optional.Optional[string] `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Result struct {
	AssessmentLog    optional.Optional[AssessmentLog]           `json:"assessment-log,omitempty" yaml:"assessment-log,omitempty"`
	Attestations     optional.Optional[[]AttestationStatements] `json:"attestations,omitempty" yaml:"attestations,omitempty"`
	Description      string                                     `json:"description" yaml:"description"`
	End              optional.Optional[time.Time]               `json:"end,omitempty" yaml:"end,omitempty"`
	Findings         optional.Optional[[]Finding]               `json:"findings,omitempty" yaml:"findings,omitempty"`
	Links            optional.Optional[[]Link]                  `json:"links,omitempty" yaml:"links,omitempty"`
	LocalDefinitions optional.Optional[LocalDefinitions]        `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Observations     optional.Optional[[]Observation]           `json:"observations,omitempty" yaml:"observations,omitempty"`
	Props            optional.Optional[[]Property]              `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          optional.Optional[string]                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ReviewedControls ReviewedControls                           `json:"reviewed-controls" yaml:"reviewed-controls"`
	Risks            optional.Optional[[]Risk]                  `json:"risks,omitempty" yaml:"risks,omitempty"`
	Start            time.Time                                  `json:"start" yaml:"start"`
	Title            string                                     `json:"title" yaml:"title"`
	UUID             string                                     `json:"uuid" yaml:"uuid"`
}

type Control struct {
	Class    optional.Optional[string]      `json:"class,omitempty" yaml:"class,omitempty"`
	Controls optional.Optional[[]Control]   `json:"controls,omitempty" yaml:"controls,omitempty"`
	ID       string                         `json:"id" yaml:"id"`
	Links    optional.Optional[[]Link]      `json:"links,omitempty" yaml:"links,omitempty"`
	Params   optional.Optional[[]Parameter] `json:"params,omitempty" yaml:"params,omitempty"`
	Parts    optional.Optional[[]Part]      `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props    optional.Optional[[]Property]  `json:"props,omitempty" yaml:"props,omitempty"`
	Title    string                         `json:"title" yaml:"title"`
}

type Group struct {
	Class    optional.Optional[string]      `json:"class,omitempty" yaml:"class,omitempty"`
	Controls optional.Optional[[]Control]   `json:"controls,omitempty" yaml:"controls,omitempty"`
	Groups   optional.Optional[[]Group]     `json:"groups,omitempty" yaml:"groups,omitempty"`
	ID       optional.Optional[string]      `json:"id,omitempty" yaml:"id,omitempty"`
	Links    optional.Optional[[]Link]      `json:"links,omitempty" yaml:"links,omitempty"`
	Params   optional.Optional[[]Parameter] `json:"params,omitempty" yaml:"params,omitempty"`
	Parts    optional.Optional[[]Part]      `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props    optional.Optional[[]Property]  `json:"props,omitempty" yaml:"props,omitempty"`
	Title    string                         `json:"title" yaml:"title"`
}

type Parameter struct {
	Class       optional.Optional[string]                `json:"class,omitempty" yaml:"class,omitempty"`
	Constraints optional.Optional[[]ParameterConstraint] `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	DependsOn   optional.Optional[string]                `json:"depends-on,omitempty" yaml:"depends-on,omitempty"`
	Guidelines  optional.Optional[[]ParameterGuideline]  `json:"guidelines,omitempty" yaml:"guidelines,omitempty"`
	ID          string                                   `json:"id" yaml:"id"`
	Label       optional.Optional[string]                `json:"label,omitempty" yaml:"label,omitempty"`
	Links       optional.Optional[[]Link]                `json:"links,omitempty" yaml:"links,omitempty"`
	Props       optional.Optional[[]Property]            `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     optional.Optional[string]                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Select      optional.Optional[ParameterSelection]    `json:"select,omitempty" yaml:"select,omitempty"`
	Usage       optional.Optional[string]                `json:"usage,omitempty" yaml:"usage,omitempty"`
	Values      optional.Optional[[]string]              `json:"values,omitempty" yaml:"values,omitempty"`
}

type Capability struct {
	ControlImplementations optional.Optional[[]ControlImplementationSet] `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Description            string                                        `json:"description" yaml:"description"`
	IncorporatesComponents optional.Optional[[]IncorporatesComponent]    `json:"incorporates-components,omitempty" yaml:"incorporates-components,omitempty"`
	Links                  optional.Optional[[]Link]                     `json:"links,omitempty" yaml:"links,omitempty"`
	Name                   string                                        `json:"name" yaml:"name"`
	Props                  optional.Optional[[]Property]                 `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks                optional.Optional[string]                     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID                   string                                        `json:"uuid" yaml:"uuid"`
}

type DefinedComponent struct {
	ControlImplementations optional.Optional[[]ControlImplementationSet] `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Description            string                                        `json:"description" yaml:"description"`
	Links                  optional.Optional[[]Link]                     `json:"links,omitempty" yaml:"links,omitempty"`
	Props                  optional.Optional[[]Property]                 `json:"props,omitempty" yaml:"props,omitempty"`
	Protocols              optional.Optional[[]Protocol]                 `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	Purpose                optional.Optional[string]                     `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Remarks                optional.Optional[string]                     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles       optional.Optional[[]ResponsibleRole]          `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Title                  string                                        `json:"title" yaml:"title"`
	Type                   string                                        `json:"type" yaml:"type"`
	UUID                   string                                        `json:"uuid" yaml:"uuid"`
}

type ImportComponentDefinition struct {
	Href string `json:"href" yaml:"href"`
}

type PlanOfActionAndMilestonesLocalDefinitions struct {
	Components     optional.Optional[[]SystemComponent] `json:"components,omitempty" yaml:"components,omitempty"`
	InventoryItems optional.Optional[[]InventoryItem]   `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	Remarks        optional.Optional[string]            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Observation struct {
	Collected        time.Time                             `json:"collected" yaml:"collected"`
	Description      string                                `json:"description" yaml:"description"`
	Expires          optional.Optional[time.Time]          `json:"expires,omitempty" yaml:"expires,omitempty"`
	Links            optional.Optional[[]Link]             `json:"links,omitempty" yaml:"links,omitempty"`
	Methods          []string                              `json:"methods" yaml:"methods"`
	Origins          optional.Optional[[]Origin]           `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props            optional.Optional[[]Property]         `json:"props,omitempty" yaml:"props,omitempty"`
	RelevantEvidence optional.Optional[[]RelevantEvidence] `json:"relevant-evidence,omitempty" yaml:"relevant-evidence,omitempty"`
	Remarks          optional.Optional[string]             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Subjects         optional.Optional[[]SubjectReference] `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	Title            optional.Optional[string]             `json:"title,omitempty" yaml:"title,omitempty"`
	Types            optional.Optional[[]string]           `json:"types,omitempty" yaml:"types,omitempty"`
	UUID             string                                `json:"uuid" yaml:"uuid"`
}

type PoamItem struct {
	Description         string                                  `json:"description" yaml:"description"`
	Links               optional.Optional[[]Link]               `json:"links,omitempty" yaml:"links,omitempty"`
	Origins             optional.Optional[[]PoamItemOrigin]     `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props               optional.Optional[[]Property]           `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedObservations optional.Optional[[]RelatedObservation] `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	RelatedRisks        optional.Optional[[]AssociatedRisk]     `json:"related-risks,omitempty" yaml:"related-risks,omitempty"`
	Remarks             optional.Optional[string]               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title               string                                  `json:"title" yaml:"title"`
	UUID                optional.Optional[string]               `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type Risk struct {
	Characterizations   optional.Optional[[]Characterization]   `json:"characterizations,omitempty" yaml:"characterizations,omitempty"`
	Deadline            optional.Optional[time.Time]            `json:"deadline,omitempty" yaml:"deadline,omitempty"`
	Description         string                                  `json:"description" yaml:"description"`
	Links               optional.Optional[[]Link]               `json:"links,omitempty" yaml:"links,omitempty"`
	MitigatingFactors   optional.Optional[[]MitigatingFactor]   `json:"mitigating-factors,omitempty" yaml:"mitigating-factors,omitempty"`
	Origins             optional.Optional[[]Origin]             `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props               optional.Optional[[]Property]           `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedObservations optional.Optional[[]RelatedObservation] `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	Remediations        optional.Optional[[]Response]           `json:"remediations,omitempty" yaml:"remediations,omitempty"`
	RiskLog             optional.Optional[RiskLog]              `json:"risk-log,omitempty" yaml:"risk-log,omitempty"`
	Statement           string                                  `json:"statement" yaml:"statement"`
	Status              string                                  `json:"status" yaml:"status"`
	ThreatIds           optional.Optional[[]ThreatId]           `json:"threat-ids,omitempty" yaml:"threat-ids,omitempty"`
	Title               string                                  `json:"title" yaml:"title"`
	UUID                string                                  `json:"uuid" yaml:"uuid"`
}

type SystemId struct {
	ID             string                    `json:"id" yaml:"id"`
	IdentifierType optional.Optional[string] `json:"identifier-type,omitempty" yaml:"identifier-type,omitempty"`
}

type Import struct {
	ExcludeControls optional.Optional[[]SelectControlById] `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
	Href            string                                 `json:"href" yaml:"href"`
	IncludeAll      optional.Optional[IncludeAll]          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls optional.Optional[[]SelectControlById] `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
}

type Merge struct {
	AsIs    optional.Optional[bool]            `json:"as-is,omitempty" yaml:"as-is,omitempty"`
	Combine optional.Optional[CombinationRule] `json:"combine,omitempty" yaml:"combine,omitempty"`
	Custom  optional.Optional[CustomGrouping]  `json:"custom,omitempty" yaml:"custom,omitempty"`
	Flat    optional.Optional[Flat]            `json:"flat,omitempty" yaml:"flat,omitempty"`
}

type Modify struct {
	Alters        optional.Optional[[]Alter]            `json:"alters,omitempty" yaml:"alters,omitempty"`
	SetParameters optional.Optional[[]ParameterSetting] `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
}

type ControlImplementation struct {
	Description             string                            `json:"description" yaml:"description"`
	ImplementedRequirements []ImplementedRequirement          `json:"implemented-requirements" yaml:"implemented-requirements"`
	SetParameters           optional.Optional[[]SetParameter] `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
}

type ImportProfile struct {
	Href    string                    `json:"href" yaml:"href"`
	Remarks optional.Optional[string] `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SystemCharacteristics struct {
	AuthorizationBoundary    AuthorizationBoundary                  `json:"authorization-boundary" yaml:"authorization-boundary"`
	DataFlow                 optional.Optional[DataFlow]            `json:"data-flow,omitempty" yaml:"data-flow,omitempty"`
	DateAuthorized           optional.Optional[string]              `json:"date-authorized,omitempty" yaml:"date-authorized,omitempty"`
	Description              string                                 `json:"description" yaml:"description"`
	Links                    optional.Optional[[]Link]              `json:"links,omitempty" yaml:"links,omitempty"`
	NetworkArchitecture      optional.Optional[NetworkArchitecture] `json:"network-architecture,omitempty" yaml:"network-architecture,omitempty"`
	Props                    optional.Optional[[]Property]          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks                  optional.Optional[string]              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties       optional.Optional[[]ResponsibleParty]  `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	SecurityImpactLevel      SecurityImpactLevel                    `json:"security-impact-level" yaml:"security-impact-level"`
	SecuritySensitivityLevel string                                 `json:"security-sensitivity-level" yaml:"security-sensitivity-level"`
	Status                   Status                                 `json:"status" yaml:"status"`
	SystemIds                []SystemId                             `json:"system-ids" yaml:"system-ids"`
	SystemInformation        SystemInformation                      `json:"system-information" yaml:"system-information"`
	SystemName               string                                 `json:"system-name" yaml:"system-name"`
	SystemNameShort          optional.Optional[string]              `json:"system-name-short,omitempty" yaml:"system-name-short,omitempty"`
}

type SystemImplementation struct {
	Components              []SystemComponent                           `json:"components" yaml:"components"`
	InventoryItems          optional.Optional[[]InventoryItem]          `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	LeveragedAuthorizations optional.Optional[[]LeveragedAuthorization] `json:"leveraged-authorizations,omitempty" yaml:"leveraged-authorizations,omitempty"`
	Links                   optional.Optional[[]Link]                   `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   optional.Optional[[]Property]               `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks                 optional.Optional[string]                   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Users                   []SystemUser                                `json:"users" yaml:"users"`
}

type AssessmentPlatform struct {
	Links          optional.Optional[[]Link]          `json:"links,omitempty" yaml:"links,omitempty"`
	Props          optional.Optional[[]Property]      `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks        optional.Optional[string]          `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title          optional.Optional[string]          `json:"title,omitempty" yaml:"title,omitempty"`
	UsesComponents optional.Optional[[]UsesComponent] `json:"uses-components,omitempty" yaml:"uses-components,omitempty"`
	UUID           string                             `json:"uuid" yaml:"uuid"`
}

type SystemComponent struct {
	Description      string                               `json:"description" yaml:"description"`
	Links            optional.Optional[[]Link]            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            optional.Optional[[]Property]        `json:"props,omitempty" yaml:"props,omitempty"`
	Protocols        optional.Optional[[]Protocol]        `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	Purpose          optional.Optional[string]            `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Remarks          optional.Optional[string]            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles optional.Optional[[]ResponsibleRole] `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Status           SystemComponentStatus                `json:"status" yaml:"status"`
	Title            string                               `json:"title" yaml:"title"`
	Type             string                               `json:"type" yaml:"type"`
	UUID             string                               `json:"uuid" yaml:"uuid"`
}

type SelectSubjectById struct {
	Links       optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     optional.Optional[string]     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	SubjectUuid string                        `json:"subject-uuid" yaml:"subject-uuid"`
	Type        string                        `json:"type" yaml:"type"`
}

type IncludeAll = map[string]interface{}

type Link struct {
	Href      string                    `json:"href" yaml:"href"`
	MediaType optional.Optional[string] `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Rel       optional.Optional[string] `json:"rel,omitempty" yaml:"rel,omitempty"`
	Text      optional.Optional[string] `json:"text,omitempty" yaml:"text,omitempty"`
}

type Property struct {
	Class   optional.Optional[string] `json:"class,omitempty" yaml:"class,omitempty"`
	Name    string                    `json:"name" yaml:"name"`
	Ns      optional.Optional[string] `json:"ns,omitempty" yaml:"ns,omitempty"`
	Remarks optional.Optional[string] `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID    optional.Optional[string] `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Value   string                    `json:"value" yaml:"value"`
}

type Resource struct {
	Base64      optional.Optional[Base64]         `json:"base64,omitempty" yaml:"base64,omitempty"`
	Citation    optional.Optional[Citation]       `json:"citation,omitempty" yaml:"citation,omitempty"`
	Description optional.Optional[string]         `json:"description,omitempty" yaml:"description,omitempty"`
	DocumentIds optional.Optional[[]DocumentId]   `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	Props       optional.Optional[[]Property]     `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     optional.Optional[string]         `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Rlinks      optional.Optional[[]ResourceLink] `json:"rlinks,omitempty" yaml:"rlinks,omitempty"`
	Title       optional.Optional[string]         `json:"title,omitempty" yaml:"title,omitempty"`
	UUID        string                            `json:"uuid" yaml:"uuid"`
}

type Activity struct {
	Description      string                               `json:"description" yaml:"description"`
	Links            optional.Optional[[]Link]            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            optional.Optional[[]Property]        `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedControls  optional.Optional[ReviewedControls]  `json:"related-controls,omitempty" yaml:"related-controls,omitempty"`
	Remarks          optional.Optional[string]            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles optional.Optional[[]ResponsibleRole] `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Steps            optional.Optional[[]Step]            `json:"steps,omitempty" yaml:"steps,omitempty"`
	Title            optional.Optional[string]            `json:"title,omitempty" yaml:"title,omitempty"`
	UUID             string                               `json:"uuid" yaml:"uuid"`
}

type InventoryItem struct {
	Description           string                                    `json:"description" yaml:"description"`
	ImplementedComponents optional.Optional[[]ImplementedComponent] `json:"implemented-components,omitempty" yaml:"implemented-components,omitempty"`
	Links                 optional.Optional[[]Link]                 `json:"links,omitempty" yaml:"links,omitempty"`
	Props                 optional.Optional[[]Property]             `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks               optional.Optional[string]                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties    optional.Optional[[]ResponsibleParty]     `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	UUID                  string                                    `json:"uuid" yaml:"uuid"`
}

type LocalObjective struct {
	ControlId   string                        `json:"control-id" yaml:"control-id"`
	Description optional.Optional[string]     `json:"description,omitempty" yaml:"description,omitempty"`
	Links       optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Parts       []Part                        `json:"parts" yaml:"parts"`
	Props       optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     optional.Optional[string]     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SystemUser struct {
	AuthorizedPrivileges optional.Optional[[]AuthorizedPrivilege] `json:"authorized-privileges,omitempty" yaml:"authorized-privileges,omitempty"`
	Description          optional.Optional[string]                `json:"description,omitempty" yaml:"description,omitempty"`
	Links                optional.Optional[[]Link]                `json:"links,omitempty" yaml:"links,omitempty"`
	Props                optional.Optional[[]Property]            `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              optional.Optional[string]                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleIds              optional.Optional[[]string]              `json:"role-ids,omitempty" yaml:"role-ids,omitempty"`
	ShortName            optional.Optional[string]                `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	Title                optional.Optional[string]                `json:"title,omitempty" yaml:"title,omitempty"`
	UUID                 string                                   `json:"uuid" yaml:"uuid"`
}

type DocumentId struct {
	Identifier string                    `json:"identifier" yaml:"identifier"`
	Scheme     optional.Optional[string] `json:"scheme,omitempty" yaml:"scheme,omitempty"`
}

type Location struct {
	Address          Address                              `json:"address" yaml:"address"`
	EmailAddresses   optional.Optional[[]string]          `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	Links            optional.Optional[[]Link]            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            optional.Optional[[]Property]        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          optional.Optional[string]            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	TelephoneNumbers optional.Optional[[]TelephoneNumber] `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	Title            optional.Optional[string]            `json:"title,omitempty" yaml:"title,omitempty"`
	Urls             optional.Optional[[]string]          `json:"urls,omitempty" yaml:"urls,omitempty"`
	UUID             string                               `json:"uuid" yaml:"uuid"`
}

type Party struct {
	Addresses             optional.Optional[[]Address]                 `json:"addresses,omitempty" yaml:"addresses,omitempty"`
	EmailAddresses        optional.Optional[[]string]                  `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	ExternalIds           optional.Optional[[]PartyExternalIdentifier] `json:"external-ids,omitempty" yaml:"external-ids,omitempty"`
	Links                 optional.Optional[[]Link]                    `json:"links,omitempty" yaml:"links,omitempty"`
	LocationUuids         optional.Optional[[]string]                  `json:"location-uuids,omitempty" yaml:"location-uuids,omitempty"`
	MemberOfOrganizations optional.Optional[[]string]                  `json:"member-of-organizations,omitempty" yaml:"member-of-organizations,omitempty"`
	Name                  optional.Optional[string]                    `json:"name,omitempty" yaml:"name,omitempty"`
	Props                 optional.Optional[[]Property]                `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks               optional.Optional[string]                    `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ShortName             optional.Optional[string]                    `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	TelephoneNumbers      optional.Optional[[]TelephoneNumber]         `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	Type                  string                                       `json:"type" yaml:"type"`
	UUID                  string                                       `json:"uuid" yaml:"uuid"`
}

type ResponsibleParty struct {
	Links      optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuids []string                      `json:"party-uuids" yaml:"party-uuids"`
	Props      optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks    optional.Optional[string]     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleId     string                        `json:"role-id" yaml:"role-id"`
}

type Revision struct {
	LastModified optional.Optional[time.Time]  `json:"last-modified,omitempty" yaml:"last-modified,omitempty"`
	Links        optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	OscalVersion optional.Optional[string]     `json:"oscal-version,omitempty" yaml:"oscal-version,omitempty"`
	Props        optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Published    optional.Optional[time.Time]  `json:"published,omitempty" yaml:"published,omitempty"`
	Remarks      optional.Optional[string]     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title        optional.Optional[string]     `json:"title,omitempty" yaml:"title,omitempty"`
	Version      string                        `json:"version" yaml:"version"`
}

type Role struct {
	Description optional.Optional[string]     `json:"description,omitempty" yaml:"description,omitempty"`
	ID          string                        `json:"id" yaml:"id"`
	Links       optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     optional.Optional[string]     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ShortName   optional.Optional[string]     `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	Title       string                        `json:"title" yaml:"title"`
}

type ReferencedControlObjectives struct {
	Description       optional.Optional[string]                `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeObjectives optional.Optional[[]SelectObjectiveById] `json:"exclude-objectives,omitempty" yaml:"exclude-objectives,omitempty"`
	IncludeAll        optional.Optional[IncludeAll]            `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeObjectives optional.Optional[[]SelectObjectiveById] `json:"include-objectives,omitempty" yaml:"include-objectives,omitempty"`
	Links             optional.Optional[[]Link]                `json:"links,omitempty" yaml:"links,omitempty"`
	Props             optional.Optional[[]Property]            `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks           optional.Optional[string]                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type AssessedControls struct {
	Description     optional.Optional[string]          `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeControls optional.Optional[[]SelectControl] `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
	IncludeAll      optional.Optional[IncludeAll]      `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls optional.Optional[[]SelectControl] `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
	Links           optional.Optional[[]Link]          `json:"links,omitempty" yaml:"links,omitempty"`
	Props           optional.Optional[[]Property]      `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks         optional.Optional[string]          `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type AssociatedActivity struct {
	ActivityUuid     string                               `json:"activity-uuid" yaml:"activity-uuid"`
	Links            optional.Optional[[]Link]            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            optional.Optional[[]Property]        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          optional.Optional[string]            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles optional.Optional[[]ResponsibleRole] `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Subjects         []AssessmentSubject                  `json:"subjects" yaml:"subjects"`
}

type TaskDependency struct {
	Remarks  optional.Optional[string] `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	TaskUuid string                    `json:"task-uuid" yaml:"task-uuid"`
}

type ResponsibleRole struct {
	Links      optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuids optional.Optional[[]string]   `json:"party-uuids,omitempty" yaml:"party-uuids,omitempty"`
	Props      optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks    optional.Optional[string]     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleId     string                        `json:"role-id" yaml:"role-id"`
}

type EventTiming struct {
	AtFrequency     optional.Optional[FrequencyCondition]   `json:"at-frequency,omitempty" yaml:"at-frequency,omitempty"`
	OnDate          optional.Optional[OnDateCondition]      `json:"on-date,omitempty" yaml:"on-date,omitempty"`
	WithinDateRange optional.Optional[OnDateRangeCondition] `json:"within-date-range,omitempty" yaml:"within-date-range,omitempty"`
}

type AssessmentPart struct {
	Class optional.Optional[string]           `json:"class,omitempty" yaml:"class,omitempty"`
	Links optional.Optional[[]Link]           `json:"links,omitempty" yaml:"links,omitempty"`
	Name  string                              `json:"name" yaml:"name"`
	Ns    optional.Optional[string]           `json:"ns,omitempty" yaml:"ns,omitempty"`
	Parts optional.Optional[[]AssessmentPart] `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props optional.Optional[[]Property]       `json:"props,omitempty" yaml:"props,omitempty"`
	Prose optional.Optional[string]           `json:"prose,omitempty" yaml:"prose,omitempty"`
	Title optional.Optional[string]           `json:"title,omitempty" yaml:"title,omitempty"`
	UUID  optional.Optional[string]           `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type AssessmentLog struct {
	Entries []AssessmentLogEntry `json:"entries" yaml:"entries"`
}

type AttestationStatements struct {
	Parts              []AssessmentPart                      `json:"parts" yaml:"parts"`
	ResponsibleParties optional.Optional[[]ResponsibleParty] `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
}

type Finding struct {
	Description                 string                                  `json:"description" yaml:"description"`
	ImplementationStatementUuid optional.Optional[string]               `json:"implementation-statement-uuid,omitempty" yaml:"implementation-statement-uuid,omitempty"`
	Links                       optional.Optional[[]Link]               `json:"links,omitempty" yaml:"links,omitempty"`
	Origins                     optional.Optional[[]Origin]             `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props                       optional.Optional[[]Property]           `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedObservations         optional.Optional[[]RelatedObservation] `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	RelatedRisks                optional.Optional[[]AssociatedRisk]     `json:"related-risks,omitempty" yaml:"related-risks,omitempty"`
	Remarks                     optional.Optional[string]               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Target                      FindingTarget                           `json:"target" yaml:"target"`
	Title                       string                                  `json:"title" yaml:"title"`
	UUID                        string                                  `json:"uuid" yaml:"uuid"`
}

type Part struct {
	Class optional.Optional[string]     `json:"class,omitempty" yaml:"class,omitempty"`
	ID    optional.Optional[string]     `json:"id,omitempty" yaml:"id,omitempty"`
	Links optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Name  string                        `json:"name" yaml:"name"`
	Ns    optional.Optional[string]     `json:"ns,omitempty" yaml:"ns,omitempty"`
	Parts optional.Optional[[]Part]     `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Prose optional.Optional[string]     `json:"prose,omitempty" yaml:"prose,omitempty"`
	Title optional.Optional[string]     `json:"title,omitempty" yaml:"title,omitempty"`
}

type ParameterConstraint struct {
	Description optional.Optional[string]           `json:"description,omitempty" yaml:"description,omitempty"`
	Tests       optional.Optional[[]ConstraintTest] `json:"tests,omitempty" yaml:"tests,omitempty"`
}

type ParameterGuideline struct {
	Prose string `json:"prose" yaml:"prose"`
}

type ParameterSelection struct {
	Choice  optional.Optional[[]string] `json:"choice,omitempty" yaml:"choice,omitempty"`
	HowMany optional.Optional[string]   `json:"how-many,omitempty" yaml:"how-many,omitempty"`
}

type ControlImplementationSet struct {
	Description             string                                        `json:"description" yaml:"description"`
	ImplementedRequirements []ImplementedRequirementControlImplementation `json:"implemented-requirements" yaml:"implemented-requirements"`
	Links                   optional.Optional[[]Link]                     `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   optional.Optional[[]Property]                 `json:"props,omitempty" yaml:"props,omitempty"`
	SetParameters           optional.Optional[[]SetParameter]             `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Source                  string                                        `json:"source" yaml:"source"`
	UUID                    string                                        `json:"uuid" yaml:"uuid"`
}

type IncorporatesComponent struct {
	ComponentUuid string `json:"component-uuid" yaml:"component-uuid"`
	Description   string `json:"description" yaml:"description"`
}

type Protocol struct {
	Name       string                         `json:"name" yaml:"name"`
	PortRanges optional.Optional[[]PortRange] `json:"port-ranges,omitempty" yaml:"port-ranges,omitempty"`
	Title      optional.Optional[string]      `json:"title,omitempty" yaml:"title,omitempty"`
	UUID       optional.Optional[string]      `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type Origin struct {
	Actors       []OriginActor                    `json:"actors" yaml:"actors"`
	RelatedTasks optional.Optional[[]RelatedTask] `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
}

type RelevantEvidence struct {
	Description string                        `json:"description" yaml:"description"`
	Href        optional.Optional[string]     `json:"href,omitempty" yaml:"href,omitempty"`
	Links       optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     optional.Optional[string]     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SubjectReference struct {
	Links       optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     optional.Optional[string]     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	SubjectUuid string                        `json:"subject-uuid" yaml:"subject-uuid"`
	Title       optional.Optional[string]     `json:"title,omitempty" yaml:"title,omitempty"`
	Type        string                        `json:"type" yaml:"type"`
}

type PoamItemOrigin struct {
	Actors []OriginActor `json:"actors" yaml:"actors"`
}

type RelatedObservation struct {
	ObservationUuid string `json:"observation-uuid" yaml:"observation-uuid"`
}

type AssociatedRisk struct {
	RiskUuid string `json:"risk-uuid" yaml:"risk-uuid"`
}

type Characterization struct {
	Facets []Facet                       `json:"facets" yaml:"facets"`
	Links  optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Origin Origin                        `json:"origin" yaml:"origin"`
	Props  optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
}

type MitigatingFactor struct {
	Description        string                                `json:"description" yaml:"description"`
	ImplementationUuid optional.Optional[string]             `json:"implementation-uuid,omitempty" yaml:"implementation-uuid,omitempty"`
	Links              optional.Optional[[]Link]             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              optional.Optional[[]Property]         `json:"props,omitempty" yaml:"props,omitempty"`
	Subjects           optional.Optional[[]SubjectReference] `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	UUID               string                                `json:"uuid" yaml:"uuid"`
}

type Response struct {
	Description    string                             `json:"description" yaml:"description"`
	Lifecycle      string                             `json:"lifecycle" yaml:"lifecycle"`
	Links          optional.Optional[[]Link]          `json:"links,omitempty" yaml:"links,omitempty"`
	Origins        optional.Optional[[]Origin]        `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props          optional.Optional[[]Property]      `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks        optional.Optional[string]          `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RequiredAssets optional.Optional[[]RequiredAsset] `json:"required-assets,omitempty" yaml:"required-assets,omitempty"`
	Tasks          optional.Optional[[]Task]          `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Title          string                             `json:"title" yaml:"title"`
	UUID           string                             `json:"uuid" yaml:"uuid"`
}

type RiskLog struct {
	Entries []RiskLogEntry `json:"entries" yaml:"entries"`
}

type ThreatId struct {
	Href   optional.Optional[string] `json:"href,omitempty" yaml:"href,omitempty"`
	ID     string                    `json:"id" yaml:"id"`
	System string                    `json:"system" yaml:"system"`
}

type SelectControlById struct {
	Matching          optional.Optional[[]MatchControlsByPattern] `json:"matching,omitempty" yaml:"matching,omitempty"`
	WithChildControls optional.Optional[string]                   `json:"with-child-controls,omitempty" yaml:"with-child-controls,omitempty"`
	WithIds           optional.Optional[[]string]                 `json:"with-ids,omitempty" yaml:"with-ids,omitempty"`
}

type CombinationRule struct {
	Method optional.Optional[string] `json:"method,omitempty" yaml:"method,omitempty"`
}

type CustomGrouping struct {
	Groups         optional.Optional[[]ControlGroup]   `json:"groups,omitempty" yaml:"groups,omitempty"`
	InsertControls optional.Optional[[]InsertControls] `json:"insert-controls,omitempty" yaml:"insert-controls,omitempty"`
}

type Flat = map[string]interface{}

type Alter struct {
	Adds      optional.Optional[[]Add]    `json:"adds,omitempty" yaml:"adds,omitempty"`
	ControlId string                      `json:"control-id" yaml:"control-id"`
	Removes   optional.Optional[[]Remove] `json:"removes,omitempty" yaml:"removes,omitempty"`
}

type ParameterSetting struct {
	Class       optional.Optional[string]                `json:"class,omitempty" yaml:"class,omitempty"`
	Constraints optional.Optional[[]ParameterConstraint] `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	DependsOn   optional.Optional[string]                `json:"depends-on,omitempty" yaml:"depends-on,omitempty"`
	Guidelines  optional.Optional[[]ParameterGuideline]  `json:"guidelines,omitempty" yaml:"guidelines,omitempty"`
	Label       optional.Optional[string]                `json:"label,omitempty" yaml:"label,omitempty"`
	Links       optional.Optional[[]Link]                `json:"links,omitempty" yaml:"links,omitempty"`
	ParamId     string                                   `json:"param-id" yaml:"param-id"`
	Props       optional.Optional[[]Property]            `json:"props,omitempty" yaml:"props,omitempty"`
	Select      optional.Optional[ParameterSelection]    `json:"select,omitempty" yaml:"select,omitempty"`
	Usage       optional.Optional[string]                `json:"usage,omitempty" yaml:"usage,omitempty"`
	Values      optional.Optional[[]string]              `json:"values,omitempty" yaml:"values,omitempty"`
}

type ImplementedRequirement struct {
	ByComponents     optional.Optional[[]ByComponent]     `json:"by-components,omitempty" yaml:"by-components,omitempty"`
	ControlId        string                               `json:"control-id" yaml:"control-id"`
	Links            optional.Optional[[]Link]            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            optional.Optional[[]Property]        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          optional.Optional[string]            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles optional.Optional[[]ResponsibleRole] `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	SetParameters    optional.Optional[[]SetParameter]    `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Statements       optional.Optional[[]Statement]       `json:"statements,omitempty" yaml:"statements,omitempty"`
	UUID             string                               `json:"uuid" yaml:"uuid"`
}

type SetParameter struct {
	ParamId string                    `json:"param-id" yaml:"param-id"`
	Remarks optional.Optional[string] `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Values  []string                  `json:"values" yaml:"values"`
}

type AuthorizationBoundary struct {
	Description string                        `json:"description" yaml:"description"`
	Diagrams    optional.Optional[[]Diagram]  `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Links       optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     optional.Optional[string]     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type DataFlow struct {
	Description string                        `json:"description" yaml:"description"`
	Diagrams    optional.Optional[[]Diagram]  `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Links       optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     optional.Optional[string]     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type NetworkArchitecture struct {
	Description string                        `json:"description" yaml:"description"`
	Diagrams    optional.Optional[[]Diagram]  `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Links       optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     optional.Optional[string]     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SecurityImpactLevel struct {
	SecurityObjectiveAvailability    string `json:"security-objective-availability" yaml:"security-objective-availability"`
	SecurityObjectiveConfidentiality string `json:"security-objective-confidentiality" yaml:"security-objective-confidentiality"`
	SecurityObjectiveIntegrity       string `json:"security-objective-integrity" yaml:"security-objective-integrity"`
}

type Status struct {
	Remarks optional.Optional[string] `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string                    `json:"state" yaml:"state"`
}

type SystemInformation struct {
	InformationTypes []InformationType             `json:"information-types" yaml:"information-types"`
	Links            optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Props            optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
}

type LeveragedAuthorization struct {
	DateAuthorized string                        `json:"date-authorized" yaml:"date-authorized"`
	Links          optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuid      string                        `json:"party-uuid" yaml:"party-uuid"`
	Props          optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks        optional.Optional[string]     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title          string                        `json:"title" yaml:"title"`
	UUID           string                        `json:"uuid" yaml:"uuid"`
}

type UsesComponent struct {
	ComponentUuid      string                                `json:"component-uuid" yaml:"component-uuid"`
	Links              optional.Optional[[]Link]             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              optional.Optional[[]Property]         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            optional.Optional[string]             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties optional.Optional[[]ResponsibleParty] `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
}

type SystemComponentStatus struct {
	Remarks optional.Optional[string] `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string                    `json:"state" yaml:"state"`
}

type Base64 struct {
	Filename  optional.Optional[string] `json:"filename,omitempty" yaml:"filename,omitempty"`
	MediaType optional.Optional[string] `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Value     string                    `json:"value" yaml:"value"`
}

type Citation struct {
	Links optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Props optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Text  string                        `json:"text" yaml:"text"`
}

type ResourceLink struct {
	Hashes    optional.Optional[[]Hash] `json:"hashes,omitempty" yaml:"hashes,omitempty"`
	Href      string                    `json:"href" yaml:"href"`
	MediaType optional.Optional[string] `json:"media-type,omitempty" yaml:"media-type,omitempty"`
}

type Step struct {
	Description      string                               `json:"description" yaml:"description"`
	Links            optional.Optional[[]Link]            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            optional.Optional[[]Property]        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          optional.Optional[string]            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles optional.Optional[[]ResponsibleRole] `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	ReviewedControls optional.Optional[ReviewedControls]  `json:"reviewed-controls,omitempty" yaml:"reviewed-controls,omitempty"`
	Title            optional.Optional[string]            `json:"title,omitempty" yaml:"title,omitempty"`
	UUID             string                               `json:"uuid" yaml:"uuid"`
}

type ImplementedComponent struct {
	ComponentUuid      string                                `json:"component-uuid" yaml:"component-uuid"`
	Links              optional.Optional[[]Link]             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              optional.Optional[[]Property]         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            optional.Optional[string]             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties optional.Optional[[]ResponsibleParty] `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
}

type AuthorizedPrivilege struct {
	Description        optional.Optional[string] `json:"description,omitempty" yaml:"description,omitempty"`
	FunctionsPerformed []string                  `json:"functions-performed" yaml:"functions-performed"`
	Title              string                    `json:"title" yaml:"title"`
}

type Address struct {
	AddrLines  optional.Optional[[]string] `json:"addr-lines,omitempty" yaml:"addr-lines,omitempty"`
	City       optional.Optional[string]   `json:"city,omitempty" yaml:"city,omitempty"`
	Country    optional.Optional[string]   `json:"country,omitempty" yaml:"country,omitempty"`
	PostalCode optional.Optional[string]   `json:"postal-code,omitempty" yaml:"postal-code,omitempty"`
	State      optional.Optional[string]   `json:"state,omitempty" yaml:"state,omitempty"`
	Type       optional.Optional[string]   `json:"type,omitempty" yaml:"type,omitempty"`
}

type TelephoneNumber struct {
	Number string                    `json:"number" yaml:"number"`
	Type   optional.Optional[string] `json:"type,omitempty" yaml:"type,omitempty"`
}

type PartyExternalIdentifier struct {
	ID     string `json:"id" yaml:"id"`
	Scheme string `json:"scheme" yaml:"scheme"`
}

type SelectObjectiveById struct {
	ObjectiveId string `json:"objective-id" yaml:"objective-id"`
}

type SelectControl struct {
	ControlId    string                      `json:"control-id" yaml:"control-id"`
	StatementIds optional.Optional[[]string] `json:"statement-ids,omitempty" yaml:"statement-ids,omitempty"`
}

type FrequencyCondition struct {
	Period int    `json:"period" yaml:"period"`
	Unit   string `json:"unit" yaml:"unit"`
}

type OnDateCondition struct {
	Date time.Time `json:"date" yaml:"date"`
}

type OnDateRangeCondition struct {
	End   time.Time `json:"end" yaml:"end"`
	Start time.Time `json:"start" yaml:"start"`
}

type AssessmentLogEntry struct {
	Description  optional.Optional[string]        `json:"description,omitempty" yaml:"description,omitempty"`
	End          optional.Optional[time.Time]     `json:"end,omitempty" yaml:"end,omitempty"`
	Links        optional.Optional[[]Link]        `json:"links,omitempty" yaml:"links,omitempty"`
	LoggedBy     optional.Optional[[]LoggedBy]    `json:"logged-by,omitempty" yaml:"logged-by,omitempty"`
	Props        optional.Optional[[]Property]    `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedTasks optional.Optional[[]RelatedTask] `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
	Remarks      optional.Optional[string]        `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Start        time.Time                        `json:"start" yaml:"start"`
	Title        optional.Optional[string]        `json:"title,omitempty" yaml:"title,omitempty"`
	UUID         string                           `json:"uuid" yaml:"uuid"`
}

type FindingTarget struct {
	Description          optional.Optional[string]               `json:"description,omitempty" yaml:"description,omitempty"`
	ImplementationStatus optional.Optional[ImplementationStatus] `json:"implementation-status,omitempty" yaml:"implementation-status,omitempty"`
	Links                optional.Optional[[]Link]               `json:"links,omitempty" yaml:"links,omitempty"`
	Props                optional.Optional[[]Property]           `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              optional.Optional[string]               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Status               ObjectiveStatus                         `json:"status" yaml:"status"`
	TargetId             string                                  `json:"target-id" yaml:"target-id"`
	Title                optional.Optional[string]               `json:"title,omitempty" yaml:"title,omitempty"`
	Type                 string                                  `json:"type" yaml:"type"`
}

type ConstraintTest struct {
	Expression string                    `json:"expression" yaml:"expression"`
	Remarks    optional.Optional[string] `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ImplementedRequirementControlImplementation struct {
	ControlId        string                                              `json:"control-id" yaml:"control-id"`
	Description      string                                              `json:"description" yaml:"description"`
	Links            optional.Optional[[]Link]                           `json:"links,omitempty" yaml:"links,omitempty"`
	Props            optional.Optional[[]Property]                       `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          optional.Optional[string]                           `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles optional.Optional[[]ResponsibleRole]                `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	SetParameters    optional.Optional[[]SetParameter]                   `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Statements       optional.Optional[[]ControlStatementImplementation] `json:"statements,omitempty" yaml:"statements,omitempty"`
	UUID             string                                              `json:"uuid" yaml:"uuid"`
}

type PortRange struct {
	End       optional.Optional[int]    `json:"end,omitempty" yaml:"end,omitempty"`
	Start     optional.Optional[int]    `json:"start,omitempty" yaml:"start,omitempty"`
	Transport optional.Optional[string] `json:"transport,omitempty" yaml:"transport,omitempty"`
}

type OriginActor struct {
	ActorUuid string                        `json:"actor-uuid" yaml:"actor-uuid"`
	Links     optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Props     optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	RoleId    optional.Optional[string]     `json:"role-id,omitempty" yaml:"role-id,omitempty"`
	Type      string                        `json:"type" yaml:"type"`
}

type RelatedTask struct {
	IdentifiedSubject  optional.Optional[IdentifiedSubject]   `json:"identified-subject,omitempty" yaml:"identified-subject,omitempty"`
	Links              optional.Optional[[]Link]              `json:"links,omitempty" yaml:"links,omitempty"`
	Props              optional.Optional[[]Property]          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            optional.Optional[string]              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties optional.Optional[[]ResponsibleParty]  `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Subjects           optional.Optional[[]AssessmentSubject] `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	TaskUuid           string                                 `json:"task-uuid" yaml:"task-uuid"`
}

type Facet struct {
	Links   optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Name    string                        `json:"name" yaml:"name"`
	Props   optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks optional.Optional[string]     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	System  string                        `json:"system" yaml:"system"`
	Value   string                        `json:"value" yaml:"value"`
}

type RequiredAsset struct {
	Description string                                `json:"description" yaml:"description"`
	Links       optional.Optional[[]Link]             `json:"links,omitempty" yaml:"links,omitempty"`
	Props       optional.Optional[[]Property]         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     optional.Optional[string]             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Subjects    optional.Optional[[]SubjectReference] `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	Title       optional.Optional[string]             `json:"title,omitempty" yaml:"title,omitempty"`
	UUID        string                                `json:"uuid" yaml:"uuid"`
}

type RiskLogEntry struct {
	Description      optional.Optional[string]                  `json:"description,omitempty" yaml:"description,omitempty"`
	End              optional.Optional[time.Time]               `json:"end,omitempty" yaml:"end,omitempty"`
	Links            optional.Optional[[]Link]                  `json:"links,omitempty" yaml:"links,omitempty"`
	LoggedBy         optional.Optional[[]LoggedBy]              `json:"logged-by,omitempty" yaml:"logged-by,omitempty"`
	Props            optional.Optional[[]Property]              `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedResponses optional.Optional[[]RiskResponseReference] `json:"related-responses,omitempty" yaml:"related-responses,omitempty"`
	Remarks          optional.Optional[string]                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Start            time.Time                                  `json:"start" yaml:"start"`
	StatusChange     optional.Optional[string]                  `json:"status-change,omitempty" yaml:"status-change,omitempty"`
	Title            optional.Optional[string]                  `json:"title,omitempty" yaml:"title,omitempty"`
	UUID             string                                     `json:"uuid" yaml:"uuid"`
}

type MatchControlsByPattern struct {
	Pattern optional.Optional[string] `json:"pattern,omitempty" yaml:"pattern,omitempty"`
}

type ControlGroup struct {
	Class          optional.Optional[string]           `json:"class,omitempty" yaml:"class,omitempty"`
	Groups         optional.Optional[[]ControlGroup]   `json:"groups,omitempty" yaml:"groups,omitempty"`
	ID             optional.Optional[string]           `json:"id,omitempty" yaml:"id,omitempty"`
	InsertControls optional.Optional[[]InsertControls] `json:"insert-controls,omitempty" yaml:"insert-controls,omitempty"`
	Links          optional.Optional[[]Link]           `json:"links,omitempty" yaml:"links,omitempty"`
	Params         optional.Optional[[]Parameter]      `json:"params,omitempty" yaml:"params,omitempty"`
	Parts          optional.Optional[[]Part]           `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props          optional.Optional[[]Property]       `json:"props,omitempty" yaml:"props,omitempty"`
	Title          string                              `json:"title" yaml:"title"`
}

type InsertControls struct {
	ExcludeControls optional.Optional[[]SelectControlById] `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
	IncludeAll      optional.Optional[IncludeAll]          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls optional.Optional[[]SelectControlById] `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
	Order           optional.Optional[string]              `json:"order,omitempty" yaml:"order,omitempty"`
}

type Add struct {
	ById     optional.Optional[string]      `json:"by-id,omitempty" yaml:"by-id,omitempty"`
	Links    optional.Optional[[]Link]      `json:"links,omitempty" yaml:"links,omitempty"`
	Params   optional.Optional[[]Parameter] `json:"params,omitempty" yaml:"params,omitempty"`
	Parts    optional.Optional[[]Part]      `json:"parts,omitempty" yaml:"parts,omitempty"`
	Position optional.Optional[string]      `json:"position,omitempty" yaml:"position,omitempty"`
	Props    optional.Optional[[]Property]  `json:"props,omitempty" yaml:"props,omitempty"`
	Title    optional.Optional[string]      `json:"title,omitempty" yaml:"title,omitempty"`
}

type Remove struct {
	ByClass    optional.Optional[string] `json:"by-class,omitempty" yaml:"by-class,omitempty"`
	ById       optional.Optional[string] `json:"by-id,omitempty" yaml:"by-id,omitempty"`
	ByItemName optional.Optional[string] `json:"by-item-name,omitempty" yaml:"by-item-name,omitempty"`
	ByName     optional.Optional[string] `json:"by-name,omitempty" yaml:"by-name,omitempty"`
	ByNs       optional.Optional[string] `json:"by-ns,omitempty" yaml:"by-ns,omitempty"`
}

type ByComponent struct {
	ComponentUuid        string                                                            `json:"component-uuid" yaml:"component-uuid"`
	Description          string                                                            `json:"description" yaml:"description"`
	Export               optional.Optional[Export]                                         `json:"export,omitempty" yaml:"export,omitempty"`
	ImplementationStatus optional.Optional[ImplementationStatus]                           `json:"implementation-status,omitempty" yaml:"implementation-status,omitempty"`
	Inherited            optional.Optional[[]InheritedControlImplementation]               `json:"inherited,omitempty" yaml:"inherited,omitempty"`
	Links                optional.Optional[[]Link]                                         `json:"links,omitempty" yaml:"links,omitempty"`
	Props                optional.Optional[[]Property]                                     `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              optional.Optional[string]                                         `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles     optional.Optional[[]ResponsibleRole]                              `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Satisfied            optional.Optional[[]SatisfiedControlImplementationResponsibility] `json:"satisfied,omitempty" yaml:"satisfied,omitempty"`
	SetParameters        optional.Optional[[]SetParameter]                                 `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	UUID                 string                                                            `json:"uuid" yaml:"uuid"`
}

type Statement struct {
	ByComponents     optional.Optional[[]ByComponent]     `json:"by-components,omitempty" yaml:"by-components,omitempty"`
	Links            optional.Optional[[]Link]            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            optional.Optional[[]Property]        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          optional.Optional[string]            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles optional.Optional[[]ResponsibleRole] `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	StatementId      string                               `json:"statement-id" yaml:"statement-id"`
	UUID             string                               `json:"uuid" yaml:"uuid"`
}

type Diagram struct {
	Caption     optional.Optional[string]     `json:"caption,omitempty" yaml:"caption,omitempty"`
	Description optional.Optional[string]     `json:"description,omitempty" yaml:"description,omitempty"`
	Links       optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     optional.Optional[string]     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID        string                        `json:"uuid" yaml:"uuid"`
}

type InformationType struct {
	AvailabilityImpact    AvailabilityImpactLevel                            `json:"availability-impact" yaml:"availability-impact"`
	Categorizations       optional.Optional[[]InformationTypeCategorization] `json:"categorizations,omitempty" yaml:"categorizations,omitempty"`
	ConfidentialityImpact ConfidentialityImpactLevel                         `json:"confidentiality-impact" yaml:"confidentiality-impact"`
	Description           string                                             `json:"description" yaml:"description"`
	IntegrityImpact       IntegrityImpactLevel                               `json:"integrity-impact" yaml:"integrity-impact"`
	Links                 optional.Optional[[]Link]                          `json:"links,omitempty" yaml:"links,omitempty"`
	Props                 optional.Optional[[]Property]                      `json:"props,omitempty" yaml:"props,omitempty"`
	Title                 string                                             `json:"title" yaml:"title"`
	UUID                  optional.Optional[string]                          `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type Hash struct {
	Algorithm string `json:"algorithm" yaml:"algorithm"`
	Value     string `json:"value" yaml:"value"`
}

type LoggedBy struct {
	PartyUuid string                    `json:"party-uuid" yaml:"party-uuid"`
	RoleId    optional.Optional[string] `json:"role-id,omitempty" yaml:"role-id,omitempty"`
}

type ImplementationStatus struct {
	Remarks optional.Optional[string] `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string                    `json:"state" yaml:"state"`
}

type ObjectiveStatus struct {
	Reason  optional.Optional[string] `json:"reason,omitempty" yaml:"reason,omitempty"`
	Remarks optional.Optional[string] `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string                    `json:"state" yaml:"state"`
}

type ControlStatementImplementation struct {
	Description      string                               `json:"description" yaml:"description"`
	Links            optional.Optional[[]Link]            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            optional.Optional[[]Property]        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          optional.Optional[string]            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles optional.Optional[[]ResponsibleRole] `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	StatementId      string                               `json:"statement-id" yaml:"statement-id"`
	UUID             string                               `json:"uuid" yaml:"uuid"`
}

type IdentifiedSubject struct {
	SubjectPlaceholderUuid string              `json:"subject-placeholder-uuid" yaml:"subject-placeholder-uuid"`
	Subjects               []AssessmentSubject `json:"subjects" yaml:"subjects"`
}

type RiskResponseReference struct {
	Links        optional.Optional[[]Link]        `json:"links,omitempty" yaml:"links,omitempty"`
	Props        optional.Optional[[]Property]    `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedTasks optional.Optional[[]RelatedTask] `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
	Remarks      optional.Optional[string]        `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponseUuid string                           `json:"response-uuid" yaml:"response-uuid"`
}

type Export struct {
	Description      optional.Optional[string]                                `json:"description,omitempty" yaml:"description,omitempty"`
	Links            optional.Optional[[]Link]                                `json:"links,omitempty" yaml:"links,omitempty"`
	Props            optional.Optional[[]Property]                            `json:"props,omitempty" yaml:"props,omitempty"`
	Provided         optional.Optional[[]ProvidedControlImplementation]       `json:"provided,omitempty" yaml:"provided,omitempty"`
	Remarks          optional.Optional[string]                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Responsibilities optional.Optional[[]ControlImplementationResponsibility] `json:"responsibilities,omitempty" yaml:"responsibilities,omitempty"`
}

type InheritedControlImplementation struct {
	Description      string                               `json:"description" yaml:"description"`
	Links            optional.Optional[[]Link]            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            optional.Optional[[]Property]        `json:"props,omitempty" yaml:"props,omitempty"`
	ProvidedUuid     optional.Optional[string]            `json:"provided-uuid,omitempty" yaml:"provided-uuid,omitempty"`
	ResponsibleRoles optional.Optional[[]ResponsibleRole] `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	UUID             string                               `json:"uuid" yaml:"uuid"`
}

type SatisfiedControlImplementationResponsibility struct {
	Description        string                               `json:"description" yaml:"description"`
	Links              optional.Optional[[]Link]            `json:"links,omitempty" yaml:"links,omitempty"`
	Props              optional.Optional[[]Property]        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            optional.Optional[string]            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibilityUuid optional.Optional[string]            `json:"responsibility-uuid,omitempty" yaml:"responsibility-uuid,omitempty"`
	ResponsibleRoles   optional.Optional[[]ResponsibleRole] `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	UUID               string                               `json:"uuid" yaml:"uuid"`
}

type AvailabilityImpactLevel struct {
	AdjustmentJustification optional.Optional[string]     `json:"adjustment-justification,omitempty" yaml:"adjustment-justification,omitempty"`
	Base                    string                        `json:"base" yaml:"base"`
	Links                   optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Selected                optional.Optional[string]     `json:"selected,omitempty" yaml:"selected,omitempty"`
}

type InformationTypeCategorization struct {
	InformationTypeIds optional.Optional[[]string] `json:"information-type-ids,omitempty" yaml:"information-type-ids,omitempty"`
	System             string                      `json:"system" yaml:"system"`
}

type ConfidentialityImpactLevel struct {
	AdjustmentJustification optional.Optional[string]     `json:"adjustment-justification,omitempty" yaml:"adjustment-justification,omitempty"`
	Base                    string                        `json:"base" yaml:"base"`
	Links                   optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Selected                optional.Optional[string]     `json:"selected,omitempty" yaml:"selected,omitempty"`
}

type IntegrityImpactLevel struct {
	AdjustmentJustification optional.Optional[string]     `json:"adjustment-justification,omitempty" yaml:"adjustment-justification,omitempty"`
	Base                    string                        `json:"base" yaml:"base"`
	Links                   optional.Optional[[]Link]     `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   optional.Optional[[]Property] `json:"props,omitempty" yaml:"props,omitempty"`
	Selected                optional.Optional[string]     `json:"selected,omitempty" yaml:"selected,omitempty"`
}

type ProvidedControlImplementation struct {
	Description      string                               `json:"description" yaml:"description"`
	Links            optional.Optional[[]Link]            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            optional.Optional[[]Property]        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          optional.Optional[string]            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles optional.Optional[[]ResponsibleRole] `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	UUID             string                               `json:"uuid" yaml:"uuid"`
}

type ControlImplementationResponsibility struct {
	Description      string                               `json:"description" yaml:"description"`
	Links            optional.Optional[[]Link]            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            optional.Optional[[]Property]        `json:"props,omitempty" yaml:"props,omitempty"`
	ProvidedUuid     optional.Optional[string]            `json:"provided-uuid,omitempty" yaml:"provided-uuid,omitempty"`
	Remarks          optional.Optional[string]            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles optional.Optional[[]ResponsibleRole] `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	UUID             string                               `json:"uuid" yaml:"uuid"`
}
