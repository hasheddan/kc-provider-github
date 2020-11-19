/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"context"

	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

// teamName extracts the partially qualified URL of a Network.
func teamName() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		n, ok := mg.(*Team)
		if !ok {
			return ""
		}
		return n.GetName()
	}
}

// ResolveReferences of this Membership
func (mg *Membership) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	// Resolve spec.forProvider.Team
	rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.Team,
		Reference:    mg.Spec.ForProvider.TeamRef,
		Selector:     mg.Spec.ForProvider.TeamSelector,
		To:           reference.To{Managed: &Team{}, List: &TeamList{}},
		Extract:      teamName(),
	})
	if err != nil {
		return errors.Wrap(err, "spec.forProvider.Team")
	}
	mg.Spec.ForProvider.Team = rsp.ResolvedValue
	mg.Spec.ForProvider.TeamRef = rsp.ResolvedReference

	return nil
}
