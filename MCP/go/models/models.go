package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// CatalogItem represents the CatalogItem schema from the OpenAPI specification
type CatalogItem struct {
	Supportedstorage interface{} `json:"SupportedStorage,omitempty"`
	Supporteduplinkgbps interface{} `json:"SupportedUplinkGbps,omitempty"`
	Weightlbs interface{} `json:"WeightLbs,omitempty"`
	Catalogitemid interface{} `json:"CatalogItemId,omitempty"`
	Ec2capacities interface{} `json:"EC2Capacities,omitempty"`
	Itemstatus interface{} `json:"ItemStatus,omitempty"`
	Powerkva interface{} `json:"PowerKva,omitempty"`
}

// ConnectionDetails represents the ConnectionDetails schema from the OpenAPI specification
type ConnectionDetails struct {
	Serverpublickey interface{} `json:"ServerPublicKey,omitempty"`
	Servertunneladdress interface{} `json:"ServerTunnelAddress,omitempty"`
	Allowedips interface{} `json:"AllowedIps,omitempty"`
	Clientpublickey interface{} `json:"ClientPublicKey,omitempty"`
	Clienttunneladdress interface{} `json:"ClientTunnelAddress,omitempty"`
	Serverendpoint interface{} `json:"ServerEndpoint,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"Tags"`
}

// ListSitesInput represents the ListSitesInput schema from the OpenAPI specification
type ListSitesInput struct {
}

// UpdateSiteOutput represents the UpdateSiteOutput schema from the OpenAPI specification
type UpdateSiteOutput struct {
	Site Site `json:"Site,omitempty"` // Information about a site.
}

// UpdateSiteRackPhysicalPropertiesInput represents the UpdateSiteRackPhysicalPropertiesInput schema from the OpenAPI specification
type UpdateSiteRackPhysicalPropertiesInput struct {
	Uplinkgbps interface{} `json:"UplinkGbps,omitempty"`
	Powerconnector interface{} `json:"PowerConnector,omitempty"`
	Powerdrawkva interface{} `json:"PowerDrawKva,omitempty"`
	Powerfeeddrop interface{} `json:"PowerFeedDrop,omitempty"`
	Powerphase interface{} `json:"PowerPhase,omitempty"`
	Uplinkcount interface{} `json:"UplinkCount,omitempty"`
	Fiberopticcabletype interface{} `json:"FiberOpticCableType,omitempty"`
	Maximumsupportedweightlbs interface{} `json:"MaximumSupportedWeightLbs,omitempty"`
	Opticalstandard interface{} `json:"OpticalStandard,omitempty"`
}

// ListCatalogItemsOutput represents the ListCatalogItemsOutput schema from the OpenAPI specification
type ListCatalogItemsOutput struct {
	Catalogitems interface{} `json:"CatalogItems,omitempty"`
	Nexttoken string `json:"NextToken,omitempty"` // The pagination token.
}

// OrderSummary represents the OrderSummary schema from the OpenAPI specification
type OrderSummary struct {
	Orderid interface{} `json:"OrderId,omitempty"`
	Ordersubmissiondate interface{} `json:"OrderSubmissionDate,omitempty"`
	Ordertype interface{} `json:"OrderType,omitempty"`
	Outpostid interface{} `json:"OutpostId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Lineitemcountsbystatus interface{} `json:"LineItemCountsByStatus,omitempty"`
	Orderfulfilleddate interface{} `json:"OrderFulfilledDate,omitempty"`
}

// CancelOrderOutput represents the CancelOrderOutput schema from the OpenAPI specification
type CancelOrderOutput struct {
}

// UpdateOutpostOutput represents the UpdateOutpostOutput schema from the OpenAPI specification
type UpdateOutpostOutput struct {
	Outpost Outpost `json:"Outpost,omitempty"` // Information about an Outpost.
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// CreateOrderInput represents the CreateOrderInput schema from the OpenAPI specification
type CreateOrderInput struct {
	Lineitems interface{} `json:"LineItems"`
	Outpostidentifier interface{} `json:"OutpostIdentifier"`
	Paymentoption interface{} `json:"PaymentOption"`
	Paymentterm interface{} `json:"PaymentTerm,omitempty"`
}

// StartConnectionRequest represents the StartConnectionRequest schema from the OpenAPI specification
type StartConnectionRequest struct {
	Networkinterfacedeviceindex interface{} `json:"NetworkInterfaceDeviceIndex"`
	Assetid interface{} `json:"AssetId"`
	Clientpublickey interface{} `json:"ClientPublicKey"`
	Deviceserialnumber interface{} `json:"DeviceSerialNumber"`
}

// ListCatalogItemsInput represents the ListCatalogItemsInput schema from the OpenAPI specification
type ListCatalogItemsInput struct {
}

// GetSiteOutput represents the GetSiteOutput schema from the OpenAPI specification
type GetSiteOutput struct {
	Site Site `json:"Site,omitempty"` // Information about a site.
}

// GetOrderInput represents the GetOrderInput schema from the OpenAPI specification
type GetOrderInput struct {
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// CancelOrderInput represents the CancelOrderInput schema from the OpenAPI specification
type CancelOrderInput struct {
}

// GetSiteAddressOutput represents the GetSiteAddressOutput schema from the OpenAPI specification
type GetSiteAddressOutput struct {
	Addresstype interface{} `json:"AddressType,omitempty"`
	Siteid string `json:"SiteId,omitempty"` // The ID of the site.
	Address interface{} `json:"Address,omitempty"`
}

// DeleteSiteInput represents the DeleteSiteInput schema from the OpenAPI specification
type DeleteSiteInput struct {
}

// Site represents the Site schema from the OpenAPI specification
type Site struct {
	Tags interface{} `json:"Tags,omitempty"`
	Operatingaddresscity interface{} `json:"OperatingAddressCity,omitempty"`
	Operatingaddressstateorregion interface{} `json:"OperatingAddressStateOrRegion,omitempty"`
	Rackphysicalproperties interface{} `json:"RackPhysicalProperties,omitempty"`
	Description string `json:"Description,omitempty"` // The description of the site.
	Name string `json:"Name,omitempty"` // The name of the site.
	Notes interface{} `json:"Notes,omitempty"`
	Operatingaddresscountrycode interface{} `json:"OperatingAddressCountryCode,omitempty"`
	Accountid string `json:"AccountId,omitempty"` // The ID of the Amazon Web Services account.
	Sitearn string `json:"SiteArn,omitempty"` // The Amazon Resource Name (ARN) of the site.
	Siteid string `json:"SiteId,omitempty"` // The ID of the site.
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// ListAssetsOutput represents the ListAssetsOutput schema from the OpenAPI specification
type ListAssetsOutput struct {
	Nexttoken string `json:"NextToken,omitempty"` // The pagination token.
	Assets interface{} `json:"Assets,omitempty"`
}

// CreateOutpostOutput represents the CreateOutpostOutput schema from the OpenAPI specification
type CreateOutpostOutput struct {
	Outpost Outpost `json:"Outpost,omitempty"` // Information about an Outpost.
}

// AssetInfo represents the AssetInfo schema from the OpenAPI specification
type AssetInfo struct {
	Rackid interface{} `json:"RackId,omitempty"`
	Assetid interface{} `json:"AssetId,omitempty"`
	Assetlocation interface{} `json:"AssetLocation,omitempty"`
	Assettype interface{} `json:"AssetType,omitempty"`
	Computeattributes interface{} `json:"ComputeAttributes,omitempty"`
}

// AssetLocation represents the AssetLocation schema from the OpenAPI specification
type AssetLocation struct {
	Rackelevation interface{} `json:"RackElevation,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// ListOutpostsOutput represents the ListOutpostsOutput schema from the OpenAPI specification
type ListOutpostsOutput struct {
	Nexttoken string `json:"NextToken,omitempty"` // The pagination token.
	Outposts []Outpost `json:"Outposts,omitempty"` // Information about the Outposts.
}

// LineItem represents the LineItem schema from the OpenAPI specification
type LineItem struct {
	Status interface{} `json:"Status,omitempty"`
	Assetinformationlist interface{} `json:"AssetInformationList,omitempty"`
	Catalogitemid interface{} `json:"CatalogItemId,omitempty"`
	Lineitemid interface{} `json:"LineItemId,omitempty"`
	Previouslineitemid interface{} `json:"PreviousLineItemId,omitempty"`
	Previousorderid interface{} `json:"PreviousOrderId,omitempty"`
	Quantity interface{} `json:"Quantity,omitempty"`
	Shipmentinformation interface{} `json:"ShipmentInformation,omitempty"`
}

// Address represents the Address schema from the OpenAPI specification
type Address struct {
	Municipality interface{} `json:"Municipality,omitempty"`
	Addressline2 interface{} `json:"AddressLine2,omitempty"`
	Contactname interface{} `json:"ContactName,omitempty"`
	Countrycode interface{} `json:"CountryCode"`
	Districtorcounty interface{} `json:"DistrictOrCounty,omitempty"`
	Postalcode interface{} `json:"PostalCode"`
	Stateorregion interface{} `json:"StateOrRegion"`
	Addressline1 interface{} `json:"AddressLine1"`
	Addressline3 interface{} `json:"AddressLine3,omitempty"`
	City interface{} `json:"City"`
	Contactphonenumber interface{} `json:"ContactPhoneNumber,omitempty"`
}

// ComputeAttributes represents the ComputeAttributes schema from the OpenAPI specification
type ComputeAttributes struct {
	Hostid interface{} `json:"HostId,omitempty"`
	State interface{} `json:"State,omitempty"`
}

// GetOutpostInstanceTypesInput represents the GetOutpostInstanceTypesInput schema from the OpenAPI specification
type GetOutpostInstanceTypesInput struct {
}

// LineItemStatusCounts represents the LineItemStatusCounts schema from the OpenAPI specification
type LineItemStatusCounts struct {
}

// ListOrdersOutput represents the ListOrdersOutput schema from the OpenAPI specification
type ListOrdersOutput struct {
	Nexttoken string `json:"NextToken,omitempty"` // The pagination token.
	Orders interface{} `json:"Orders,omitempty"`
}

// ShipmentInformation represents the ShipmentInformation schema from the OpenAPI specification
type ShipmentInformation struct {
	Shipmentcarrier interface{} `json:"ShipmentCarrier,omitempty"`
	Shipmenttrackingnumber interface{} `json:"ShipmentTrackingNumber,omitempty"`
}

// GetCatalogItemInput represents the GetCatalogItemInput schema from the OpenAPI specification
type GetCatalogItemInput struct {
}

// UpdateOutpostInput represents the UpdateOutpostInput schema from the OpenAPI specification
type UpdateOutpostInput struct {
	Description string `json:"Description,omitempty"` // The description of the Outpost.
	Name string `json:"Name,omitempty"` // The name of the Outpost.
	Supportedhardwaretype interface{} `json:"SupportedHardwareType,omitempty"`
}

// EC2Capacity represents the EC2Capacity schema from the OpenAPI specification
type EC2Capacity struct {
	Family interface{} `json:"Family,omitempty"`
	Maxsize interface{} `json:"MaxSize,omitempty"`
	Quantity interface{} `json:"Quantity,omitempty"`
}

// UpdateSiteAddressInput represents the UpdateSiteAddressInput schema from the OpenAPI specification
type UpdateSiteAddressInput struct {
	Address interface{} `json:"Address"`
	Addresstype interface{} `json:"AddressType"`
}

// GetConnectionRequest represents the GetConnectionRequest schema from the OpenAPI specification
type GetConnectionRequest struct {
}

// ListOrdersInput represents the ListOrdersInput schema from the OpenAPI specification
type ListOrdersInput struct {
}

// GetOrderOutput represents the GetOrderOutput schema from the OpenAPI specification
type GetOrderOutput struct {
	Order Order `json:"Order,omitempty"` // Information about an order.
}

// GetSiteAddressInput represents the GetSiteAddressInput schema from the OpenAPI specification
type GetSiteAddressInput struct {
}

// ListAssetsInput represents the ListAssetsInput schema from the OpenAPI specification
type ListAssetsInput struct {
}

// RackPhysicalProperties represents the RackPhysicalProperties schema from the OpenAPI specification
type RackPhysicalProperties struct {
	Fiberopticcabletype interface{} `json:"FiberOpticCableType,omitempty"`
	Maximumsupportedweightlbs interface{} `json:"MaximumSupportedWeightLbs,omitempty"`
	Powerdrawkva interface{} `json:"PowerDrawKva,omitempty"`
	Powerconnector interface{} `json:"PowerConnector,omitempty"`
	Opticalstandard interface{} `json:"OpticalStandard,omitempty"`
	Powerphase interface{} `json:"PowerPhase,omitempty"`
	Uplinkgbps interface{} `json:"UplinkGbps,omitempty"`
	Powerfeeddrop interface{} `json:"PowerFeedDrop,omitempty"`
	Uplinkcount interface{} `json:"UplinkCount,omitempty"`
}

// GetSiteInput represents the GetSiteInput schema from the OpenAPI specification
type GetSiteInput struct {
}

// Outpost represents the Outpost schema from the OpenAPI specification
type Outpost struct {
	Ownerid string `json:"OwnerId,omitempty"` // The Amazon Web Services account ID of the Outpost owner.
	Supportedhardwaretype interface{} `json:"SupportedHardwareType,omitempty"`
	Outpostarn string `json:"OutpostArn,omitempty"` // The Amazon Resource Name (ARN) of the Outpost.
	Sitearn string `json:"SiteArn,omitempty"` // The Amazon Resource Name (ARN) of the site.
	Siteid string `json:"SiteId,omitempty"` // The ID of the site.
	Availabilityzone string `json:"AvailabilityZone,omitempty"` // The Availability Zone.
	Outpostid interface{} `json:"OutpostId,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Availabilityzoneid string `json:"AvailabilityZoneId,omitempty"` // The ID of the Availability Zone.
	Lifecyclestatus string `json:"LifeCycleStatus,omitempty"` // The life cycle status.
	Name string `json:"Name,omitempty"` // The name of the Outpost.
	Description string `json:"Description,omitempty"` // The description of the Outpost.
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// StartConnectionResponse represents the StartConnectionResponse schema from the OpenAPI specification
type StartConnectionResponse struct {
	Connectionid interface{} `json:"ConnectionId,omitempty"`
	Underlayipaddress interface{} `json:"UnderlayIpAddress,omitempty"`
}

// Order represents the Order schema from the OpenAPI specification
type Order struct {
	Orderid interface{} `json:"OrderId,omitempty"`
	Ordersubmissiondate interface{} `json:"OrderSubmissionDate,omitempty"`
	Ordertype interface{} `json:"OrderType,omitempty"`
	Paymentoption interface{} `json:"PaymentOption,omitempty"`
	Outpostid interface{} `json:"OutpostId,omitempty"`
	Paymentterm interface{} `json:"PaymentTerm,omitempty"`
	Lineitems interface{} `json:"LineItems,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Orderfulfilleddate interface{} `json:"OrderFulfilledDate,omitempty"`
}

// CreateOutpostInput represents the CreateOutpostInput schema from the OpenAPI specification
type CreateOutpostInput struct {
	Availabilityzoneid string `json:"AvailabilityZoneId,omitempty"` // The ID of the Availability Zone.
	Description string `json:"Description,omitempty"` // The description of the Outpost.
	Name string `json:"Name"` // The name of the Outpost.
	Siteid interface{} `json:"SiteId"`
	Supportedhardwaretype interface{} `json:"SupportedHardwareType,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Availabilityzone string `json:"AvailabilityZone,omitempty"` // The Availability Zone.
}

// GetOutpostInstanceTypesOutput represents the GetOutpostInstanceTypesOutput schema from the OpenAPI specification
type GetOutpostInstanceTypesOutput struct {
	Outpostarn string `json:"OutpostArn,omitempty"` // The Amazon Resource Name (ARN) of the Outpost.
	Outpostid interface{} `json:"OutpostId,omitempty"`
	Instancetypes []InstanceTypeItem `json:"InstanceTypes,omitempty"` // Information about the instance types.
	Nexttoken string `json:"NextToken,omitempty"` // The pagination token.
}

// GetOutpostOutput represents the GetOutpostOutput schema from the OpenAPI specification
type GetOutpostOutput struct {
	Outpost Outpost `json:"Outpost,omitempty"` // Information about an Outpost.
}

// ListSitesOutput represents the ListSitesOutput schema from the OpenAPI specification
type ListSitesOutput struct {
	Nexttoken string `json:"NextToken,omitempty"` // The pagination token.
	Sites []Site `json:"Sites,omitempty"` // Information about the sites.
}

// DeleteOutpostInput represents the DeleteOutpostInput schema from the OpenAPI specification
type DeleteOutpostInput struct {
}

// DeleteSiteOutput represents the DeleteSiteOutput schema from the OpenAPI specification
type DeleteSiteOutput struct {
}

// ListOutpostsInput represents the ListOutpostsInput schema from the OpenAPI specification
type ListOutpostsInput struct {
}

// UpdateSiteInput represents the UpdateSiteInput schema from the OpenAPI specification
type UpdateSiteInput struct {
	Description string `json:"Description,omitempty"` // The description of the site.
	Name string `json:"Name,omitempty"` // The name of the site.
	Notes interface{} `json:"Notes,omitempty"`
}

// UpdateSiteRackPhysicalPropertiesOutput represents the UpdateSiteRackPhysicalPropertiesOutput schema from the OpenAPI specification
type UpdateSiteRackPhysicalPropertiesOutput struct {
	Site Site `json:"Site,omitempty"` // Information about a site.
}

// GetCatalogItemOutput represents the GetCatalogItemOutput schema from the OpenAPI specification
type GetCatalogItemOutput struct {
	Catalogitem interface{} `json:"CatalogItem,omitempty"`
}

// LineItemAssetInformation represents the LineItemAssetInformation schema from the OpenAPI specification
type LineItemAssetInformation struct {
	Macaddresslist interface{} `json:"MacAddressList,omitempty"`
	Assetid interface{} `json:"AssetId,omitempty"`
}

// CreateSiteOutput represents the CreateSiteOutput schema from the OpenAPI specification
type CreateSiteOutput struct {
	Site Site `json:"Site,omitempty"` // Information about a site.
}

// GetOutpostInput represents the GetOutpostInput schema from the OpenAPI specification
type GetOutpostInput struct {
}

// LineItemRequest represents the LineItemRequest schema from the OpenAPI specification
type LineItemRequest struct {
	Quantity interface{} `json:"Quantity,omitempty"`
	Catalogitemid interface{} `json:"CatalogItemId,omitempty"`
}

// UpdateSiteAddressOutput represents the UpdateSiteAddressOutput schema from the OpenAPI specification
type UpdateSiteAddressOutput struct {
	Address interface{} `json:"Address,omitempty"`
	Addresstype interface{} `json:"AddressType,omitempty"`
}

// CreateOrderOutput represents the CreateOrderOutput schema from the OpenAPI specification
type CreateOrderOutput struct {
	Order interface{} `json:"Order,omitempty"`
}

// GetConnectionResponse represents the GetConnectionResponse schema from the OpenAPI specification
type GetConnectionResponse struct {
	Connectiondetails interface{} `json:"ConnectionDetails,omitempty"`
	Connectionid interface{} `json:"ConnectionId,omitempty"`
}

// InstanceTypeItem represents the InstanceTypeItem schema from the OpenAPI specification
type InstanceTypeItem struct {
	Instancetype string `json:"InstanceType,omitempty"` // The instance type.
}

// DeleteOutpostOutput represents the DeleteOutpostOutput schema from the OpenAPI specification
type DeleteOutpostOutput struct {
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// CreateSiteInput represents the CreateSiteInput schema from the OpenAPI specification
type CreateSiteInput struct {
	Shippingaddress interface{} `json:"ShippingAddress,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Description string `json:"Description,omitempty"` // The description of the site.
	Name string `json:"Name"` // The name of the site.
	Notes interface{} `json:"Notes,omitempty"`
	Operatingaddress interface{} `json:"OperatingAddress,omitempty"`
	Rackphysicalproperties interface{} `json:"RackPhysicalProperties,omitempty"`
}
