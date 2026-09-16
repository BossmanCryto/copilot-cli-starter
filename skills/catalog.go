package skills

// Skill represents an AI agent skill
type Skill struct {
	Name        string
	Description string
	Usage       string
}

// Catalog manages available skills
type Catalog struct {
	Skills map[string]*Skill
}

// NewCatalog creates a new skill catalog
func NewCatalog() *Catalog {
	return &Catalog{
		Skills: make(map[string]*Skill),
	}
}

// Register adds a skill to the catalog
func (c *Catalog) Register(skill *Skill) {
	c.Skills[skill.Name] = skill
}

// Get retrieves a skill by name
func (c *Catalog) Get(name string) *Skill {
	return c.Skills[name]
}

// List returns all available skills
func (c *Catalog) List() []*Skill {
	var skills []*Skill
	for _, skill := range c.Skills {
		skills = append(skills, skill)
	}
	return skills
}
