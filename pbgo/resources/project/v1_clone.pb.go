// Code generated by protoc-gen-clone. DO NOT EDIT.

package project_pb

func (z *V1) Clone() *V1 {
	if z == nil {
		return nil
	}
	zz := &V1{}
	zz.Id = z.Id
	zz.CreatedAt = z.CreatedAt
	zz.UpdatedAt = z.UpdatedAt
	zz.OrgID = z.OrgID
	zz.CreatedByUserID = z.CreatedByUserID
	zz.Name = z.Name
	zz.Remarks = z.Remarks
	zz0 := make(map[string]string)
	for k, v := range z.Environments {
		zz0[k] = v
	}
	zz.Environments = zz0
	zz.Source = z.Source.Clone()
	return zz
}
