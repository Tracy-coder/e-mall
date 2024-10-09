// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Tracy-coder/e-mall/data/ent/address"
	"github.com/Tracy-coder/e-mall/data/ent/carousel"
	"github.com/Tracy-coder/e-mall/data/ent/cart"
	"github.com/Tracy-coder/e-mall/data/ent/category"
	"github.com/Tracy-coder/e-mall/data/ent/email"
	"github.com/Tracy-coder/e-mall/data/ent/favourite"
	"github.com/Tracy-coder/e-mall/data/ent/notice"
	"github.com/Tracy-coder/e-mall/data/ent/order"
	"github.com/Tracy-coder/e-mall/data/ent/product"
	"github.com/Tracy-coder/e-mall/data/ent/productimg"
	"github.com/Tracy-coder/e-mall/data/ent/schema"
	"github.com/Tracy-coder/e-mall/data/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	addressMixin := schema.Address{}.Mixin()
	addressMixinFields0 := addressMixin[0].Fields()
	_ = addressMixinFields0
	addressFields := schema.Address{}.Fields()
	_ = addressFields
	// addressDescCreatedAt is the schema descriptor for created_at field.
	addressDescCreatedAt := addressMixinFields0[1].Descriptor()
	// address.DefaultCreatedAt holds the default value on creation for the created_at field.
	address.DefaultCreatedAt = addressDescCreatedAt.Default.(func() time.Time)
	// addressDescUpdatedAt is the schema descriptor for updated_at field.
	addressDescUpdatedAt := addressMixinFields0[2].Descriptor()
	// address.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	address.DefaultUpdatedAt = addressDescUpdatedAt.Default.(func() time.Time)
	// address.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	address.UpdateDefaultUpdatedAt = addressDescUpdatedAt.UpdateDefault.(func() time.Time)
	carouselMixin := schema.Carousel{}.Mixin()
	carouselMixinFields0 := carouselMixin[0].Fields()
	_ = carouselMixinFields0
	carouselFields := schema.Carousel{}.Fields()
	_ = carouselFields
	// carouselDescCreatedAt is the schema descriptor for created_at field.
	carouselDescCreatedAt := carouselMixinFields0[1].Descriptor()
	// carousel.DefaultCreatedAt holds the default value on creation for the created_at field.
	carousel.DefaultCreatedAt = carouselDescCreatedAt.Default.(func() time.Time)
	// carouselDescUpdatedAt is the schema descriptor for updated_at field.
	carouselDescUpdatedAt := carouselMixinFields0[2].Descriptor()
	// carousel.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	carousel.DefaultUpdatedAt = carouselDescUpdatedAt.Default.(func() time.Time)
	// carousel.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	carousel.UpdateDefaultUpdatedAt = carouselDescUpdatedAt.UpdateDefault.(func() time.Time)
	cartMixin := schema.Cart{}.Mixin()
	cartMixinFields0 := cartMixin[0].Fields()
	_ = cartMixinFields0
	cartFields := schema.Cart{}.Fields()
	_ = cartFields
	// cartDescCreatedAt is the schema descriptor for created_at field.
	cartDescCreatedAt := cartMixinFields0[1].Descriptor()
	// cart.DefaultCreatedAt holds the default value on creation for the created_at field.
	cart.DefaultCreatedAt = cartDescCreatedAt.Default.(func() time.Time)
	// cartDescUpdatedAt is the schema descriptor for updated_at field.
	cartDescUpdatedAt := cartMixinFields0[2].Descriptor()
	// cart.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	cart.DefaultUpdatedAt = cartDescUpdatedAt.Default.(func() time.Time)
	// cart.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	cart.UpdateDefaultUpdatedAt = cartDescUpdatedAt.UpdateDefault.(func() time.Time)
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescCreatedAt is the schema descriptor for created_at field.
	categoryDescCreatedAt := categoryFields[2].Descriptor()
	// category.DefaultCreatedAt holds the default value on creation for the created_at field.
	category.DefaultCreatedAt = categoryDescCreatedAt.Default.(func() time.Time)
	// categoryDescUpdatedAt is the schema descriptor for updated_at field.
	categoryDescUpdatedAt := categoryFields[3].Descriptor()
	// category.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	category.DefaultUpdatedAt = categoryDescUpdatedAt.Default.(func() time.Time)
	// category.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	category.UpdateDefaultUpdatedAt = categoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	emailMixin := schema.Email{}.Mixin()
	emailMixinFields0 := emailMixin[0].Fields()
	_ = emailMixinFields0
	emailMixinFields1 := emailMixin[1].Fields()
	_ = emailMixinFields1
	emailFields := schema.Email{}.Fields()
	_ = emailFields
	// emailDescCreatedAt is the schema descriptor for created_at field.
	emailDescCreatedAt := emailMixinFields0[1].Descriptor()
	// email.DefaultCreatedAt holds the default value on creation for the created_at field.
	email.DefaultCreatedAt = emailDescCreatedAt.Default.(func() time.Time)
	// emailDescUpdatedAt is the schema descriptor for updated_at field.
	emailDescUpdatedAt := emailMixinFields0[2].Descriptor()
	// email.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	email.DefaultUpdatedAt = emailDescUpdatedAt.Default.(func() time.Time)
	// email.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	email.UpdateDefaultUpdatedAt = emailDescUpdatedAt.UpdateDefault.(func() time.Time)
	// emailDescStatus is the schema descriptor for status field.
	emailDescStatus := emailMixinFields1[0].Descriptor()
	// email.DefaultStatus holds the default value on creation for the status field.
	email.DefaultStatus = emailDescStatus.Default.(uint8)
	// emailDescIsVerified is the schema descriptor for is_verified field.
	emailDescIsVerified := emailFields[1].Descriptor()
	// email.DefaultIsVerified holds the default value on creation for the is_verified field.
	email.DefaultIsVerified = emailDescIsVerified.Default.(bool)
	favouriteMixin := schema.Favourite{}.Mixin()
	favouriteMixinFields0 := favouriteMixin[0].Fields()
	_ = favouriteMixinFields0
	favouriteFields := schema.Favourite{}.Fields()
	_ = favouriteFields
	// favouriteDescCreatedAt is the schema descriptor for created_at field.
	favouriteDescCreatedAt := favouriteMixinFields0[1].Descriptor()
	// favourite.DefaultCreatedAt holds the default value on creation for the created_at field.
	favourite.DefaultCreatedAt = favouriteDescCreatedAt.Default.(func() time.Time)
	// favouriteDescUpdatedAt is the schema descriptor for updated_at field.
	favouriteDescUpdatedAt := favouriteMixinFields0[2].Descriptor()
	// favourite.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	favourite.DefaultUpdatedAt = favouriteDescUpdatedAt.Default.(func() time.Time)
	// favourite.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	favourite.UpdateDefaultUpdatedAt = favouriteDescUpdatedAt.UpdateDefault.(func() time.Time)
	noticeMixin := schema.Notice{}.Mixin()
	noticeMixinFields0 := noticeMixin[0].Fields()
	_ = noticeMixinFields0
	noticeFields := schema.Notice{}.Fields()
	_ = noticeFields
	// noticeDescCreatedAt is the schema descriptor for created_at field.
	noticeDescCreatedAt := noticeMixinFields0[1].Descriptor()
	// notice.DefaultCreatedAt holds the default value on creation for the created_at field.
	notice.DefaultCreatedAt = noticeDescCreatedAt.Default.(func() time.Time)
	// noticeDescUpdatedAt is the schema descriptor for updated_at field.
	noticeDescUpdatedAt := noticeMixinFields0[2].Descriptor()
	// notice.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notice.DefaultUpdatedAt = noticeDescUpdatedAt.Default.(func() time.Time)
	// notice.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notice.UpdateDefaultUpdatedAt = noticeDescUpdatedAt.UpdateDefault.(func() time.Time)
	orderMixin := schema.Order{}.Mixin()
	orderMixinFields0 := orderMixin[0].Fields()
	_ = orderMixinFields0
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescCreatedAt is the schema descriptor for created_at field.
	orderDescCreatedAt := orderMixinFields0[1].Descriptor()
	// order.DefaultCreatedAt holds the default value on creation for the created_at field.
	order.DefaultCreatedAt = orderDescCreatedAt.Default.(func() time.Time)
	// orderDescUpdatedAt is the schema descriptor for updated_at field.
	orderDescUpdatedAt := orderMixinFields0[2].Descriptor()
	// order.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	order.DefaultUpdatedAt = orderDescUpdatedAt.Default.(func() time.Time)
	// order.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	order.UpdateDefaultUpdatedAt = orderDescUpdatedAt.UpdateDefault.(func() time.Time)
	productMixin := schema.Product{}.Mixin()
	productMixinFields0 := productMixin[0].Fields()
	_ = productMixinFields0
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescCreatedAt is the schema descriptor for created_at field.
	productDescCreatedAt := productMixinFields0[1].Descriptor()
	// product.DefaultCreatedAt holds the default value on creation for the created_at field.
	product.DefaultCreatedAt = productDescCreatedAt.Default.(func() time.Time)
	// productDescUpdatedAt is the schema descriptor for updated_at field.
	productDescUpdatedAt := productMixinFields0[2].Descriptor()
	// product.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	product.DefaultUpdatedAt = productDescUpdatedAt.Default.(func() time.Time)
	// product.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	product.UpdateDefaultUpdatedAt = productDescUpdatedAt.UpdateDefault.(func() time.Time)
	productimgMixin := schema.ProductImg{}.Mixin()
	productimgMixinFields0 := productimgMixin[0].Fields()
	_ = productimgMixinFields0
	productimgFields := schema.ProductImg{}.Fields()
	_ = productimgFields
	// productimgDescCreatedAt is the schema descriptor for created_at field.
	productimgDescCreatedAt := productimgMixinFields0[1].Descriptor()
	// productimg.DefaultCreatedAt holds the default value on creation for the created_at field.
	productimg.DefaultCreatedAt = productimgDescCreatedAt.Default.(func() time.Time)
	// productimgDescUpdatedAt is the schema descriptor for updated_at field.
	productimgDescUpdatedAt := productimgMixinFields0[2].Descriptor()
	// productimg.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	productimg.DefaultUpdatedAt = productimgDescUpdatedAt.Default.(func() time.Time)
	// productimg.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	productimg.UpdateDefaultUpdatedAt = productimgDescUpdatedAt.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userMixinFields1[0].Descriptor()
	// user.DefaultStatus holds the default value on creation for the status field.
	user.DefaultStatus = userDescStatus.Default.(uint8)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[4].Descriptor()
	// user.DefaultAvatar holds the default value on creation for the avatar field.
	user.DefaultAvatar = userDescAvatar.Default.(string)
}
